package service

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pengdst/golang-file-upload/config"
	"github.com/pengdst/golang-file-upload/exception"
	"github.com/pengdst/golang-file-upload/model"
	"github.com/pengdst/golang-file-upload/model/claims"
	"github.com/pengdst/golang-file-upload/model/entity"
	"github.com/pengdst/golang-file-upload/repository"
	"github.com/pengdst/golang-file-upload/utils"
	"math/rand"
	"strconv"
	"time"
)

type AuthService interface {
	Login(ctx *gin.Context, payload model.LoginPayload) *model.LoginData
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

func (a *AuthServiceImpl) Login(ctx *gin.Context, payload model.LoginPayload) *model.LoginData {
	user, err := a.UserRepository.Verify(ctx, payload.Email, payload.Password)
	if err != nil {
		panic(exception.NewUnauthorizedError("wrong email or password"))
		return nil
	}

	accessToken, errr := a.GenerateAccessToken(user)
	if errr != nil {
		panic(errr)
	}

	refreshToken, errr := a.GenerateRefreshToken(user)
	if errr != nil {
		panic(errr)
	}

	return &model.LoginData{
		User: model.User{
			Name:       user.Name,
			Email:      user.Email,
			Role:       user.Role,
			CreatedAt:  user.CreatedAt,
			UpdaetedAt: user.UpdatedAt,
		},
		Token: model.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
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

func (a *AuthServiceImpl) GenerateAccessToken(user *entity.User) (string, error) {
	return jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims.AccessTokenClaims{
			UserID:    int(user.ID),
			TokenType: claims.AccessToken,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		},
	).SignedString([]byte(a.Env.JwtSecret))
}

func (a *AuthServiceImpl) GenerateRefreshToken(user *entity.User) (string, error) {
	cusKey, err := a.GenerateCustomKey(user.ID, user.TokenHash)
	if err != nil {
		return "", err
	}

	return jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims.RefreshTokenClaims{
			UserID:    int(user.ID),
			CusKey:    cusKey,
			TokenType: claims.RefreshToken,
			StandardClaims: jwt.StandardClaims{
				IssuedAt: time.Now().Unix(),
			},
		},
	).SignedString([]byte(a.Env.JwtSecret))
}

func (a *AuthServiceImpl) GenerateCustomKey(id uint, tokenHash string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(fmt.Sprintf("%v:%s", id, tokenHash)))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
