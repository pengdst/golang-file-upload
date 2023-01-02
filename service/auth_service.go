package service

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
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
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"time"
)

type AuthService interface {
	Login(ctx *gin.Context, payload model.LoginPayload) *model.LoginData
	Register(ctx *gin.Context, payload model.RegisterPayload)
	RefreshToken(ctx *gin.Context) string
	ValidateRefreshToken(tokenString string) (*claims.RefreshTokenClaims, error)
	ValidateAccessToken(tokenString string) (*claims.AccessTokenClaims, error)
	GenerateCustomKey(id uint, tokenHash string) (string, error)
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
	refreshToken, errr := a.GenerateRefreshToken(user)
	if errr != nil {
		panic(exception.NewUnauthorizedError("failed to authenticate user"))
	}

	return &model.LoginData{
		User: model.User{
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Token: model.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (a *AuthServiceImpl) RefreshToken(ctx *gin.Context) string {
	user := ctx.Request.Context().Value(model.User{}).(entity.User)
	accessToken, errr := a.GenerateAccessToken(&user)
	if errr != nil {
		panic(exception.NewUnauthorizedError("failed to authenticate user"))
	}

	return accessToken
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

func (a *AuthServiceImpl) ValidateRefreshToken(tokenString string) (*claims.RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims.RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Error("Unexpected signing method in auth token")
			return nil, errors.New("unexpected signing method in auth token")
		}

		return []byte(a.Env.JwtSecret), nil
	})

	if err != nil {
		log.Error("unable to parse claims", "error", err)
		return nil, err
	}

	tokenClaims, ok := token.Claims.(*claims.RefreshTokenClaims)
	if !ok || !token.Valid || tokenClaims.TokenType != claims.RefreshToken {
		log.Error("could not extract claims from token")
		return nil, errors.New("invalid token: authentication failed")
	}

	return tokenClaims, nil
}

func (a *AuthServiceImpl) ValidateAccessToken(tokenString string) (*claims.AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Error("Unexpected signing method in auth token")
			return nil, errors.New("unexpected signing method in auth token")
		}

		return []byte(a.Env.JwtSecret), nil
	})

	if err != nil {
		log.Error("unable to parse claims:", err)
		return nil, err
	}

	tokenClaims, ok := token.Claims.(*claims.AccessTokenClaims)
	if !ok || !token.Valid || tokenClaims.TokenType != claims.AccessToken {
		log.Error("could not extract claims from token")
		return nil, errors.New("invalid token: authentication failed")
	}

	return tokenClaims, nil
}

func (a *AuthServiceImpl) GenerateCustomKey(id uint, tokenHash string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(fmt.Sprintf("%v:%s", id, tokenHash)))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
