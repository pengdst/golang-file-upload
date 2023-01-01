package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pengdst/golang-file-upload/config"
	"github.com/pengdst/golang-file-upload/exception"
	"github.com/pengdst/golang-file-upload/model"
	"github.com/pengdst/golang-file-upload/model/entity"
	"github.com/pengdst/golang-file-upload/repository"
	"github.com/pengdst/golang-file-upload/utils"
	"math/rand"
	"strconv"
)

type AuthService interface {
	Login(ctx *gin.Context, payload model.LoginPayload) *model.User
	Register(ctx *gin.Context, payload model.RegisterPayload)
}

type AuthServiceImpl struct {
	Env            *config.EnvConf
	UserRepository repository.UserRepository
}

func NewAuthService(env *config.EnvConf, userRepository repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		Env:            env,
		UserRepository: userRepository,
	}
}

func (a *AuthServiceImpl) Login(ctx *gin.Context, payload model.LoginPayload) *model.User {
	user, err := a.UserRepository.Verify(ctx, payload.Email, payload.Password)
	if err != nil {
		panic(exception.NewUnauthorizedError("wrong email or password"))
		return nil
	}

	return &model.User{
		Name:       user.Name,
		Email:      user.Email,
		Token:      "",
		CreatedAt:  user.CreatedAt,
		UpdaetedAt: user.UpdatedAt,
	}
}

func (a *AuthServiceImpl) Register(ctx *gin.Context, payload model.RegisterPayload) {
	hashPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		panic(err)
	}

	_, err = a.UserRepository.Create(ctx, entity.User{
		Name:      payload.Name,
		Email:     payload.Email,
		TokenHash: strconv.FormatUint(rand.Uint64(), 19),
		Password:  hashPassword,
	})
	if err != nil {
		panic(exception.NewUnauthorizedError("user already exists"))
	}
}
