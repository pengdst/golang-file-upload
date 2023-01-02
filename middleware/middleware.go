package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pengdst/golang-file-upload/exception"
	"github.com/pengdst/golang-file-upload/model"
	"github.com/pengdst/golang-file-upload/repository"
	"github.com/pengdst/golang-file-upload/service"
	log "github.com/sirupsen/logrus"
	"strings"
)

type ApiMiddleware struct {
	AuthService    service.AuthService
	UserRepository repository.UserRepository
}

func NewApiMiddleware(authService service.AuthService, userRepo repository.UserRepository) *ApiMiddleware {
	return &ApiMiddleware{
		AuthService:    authService,
		UserRepository: userRepo,
	}
}

func (m *ApiMiddleware) ValidateAccessToken(ctx *gin.Context) {
	token, err := extractToken(ctx)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	userID, err := m.AuthService.ValidateAccessToken(token)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	log.Info("access token validated:", userID)

	ctx.Next()
}

func (m *ApiMiddleware) ValidateRefreshToken(ctx *gin.Context) {
	token, err := extractToken(ctx)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	tokenClaims, err := m.AuthService.ValidateRefreshToken(token)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	user, err := m.UserRepository.Get(ctx, tokenClaims.UserID)
	if err != nil {
		panic(exception.NewUnauthorizedError("cannot authorize user"))
	}

	actualCustomKey, err := m.AuthService.GenerateCustomKey(user.ID, user.TokenHash)
	if err != nil {
		panic(exception.NewUnauthorizedError("authentication failed: actual token"))
	}

	if tokenClaims.CusKey != actualCustomKey {
		panic(exception.NewUnauthorizedError("authentication failed: invalid token"))
	}

	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), model.User{}, *user))
	ctx.Next()
}

func extractToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")

	if !strings.Contains(authHeader, "Bearer ") {
		return "", errors.New("token not provided or malformed")
	}

	return strings.Replace(authHeader, "Bearer ", "", -1), nil
}
