package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pengdst/golang-file-upload/model"
	"github.com/pengdst/golang-file-upload/model/entity"
	"github.com/pengdst/golang-file-upload/repository"
	"github.com/pengdst/golang-file-upload/utils"
)

type AuthService interface {
	Login(ctx *gin.Context, payload model.LoginPayload) *model.User
	Register(ctx *gin.Context, payload model.RegisterPayload)
}

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
}

func (a *AuthServiceImpl) Login(ctx *gin.Context, payload model.LoginPayload) *model.User {
	user, err := a.UserRepository.Verify(ctx, payload.Email, payload.Password)
	if err != nil {
		//TODO unauthorized error
		panic(err)
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
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashPassword,
	})
	if err != nil {
		//TODO unauthorized error
		panic(err)
		return
	}
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
	}
}
