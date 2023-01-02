package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pengdst/golang-file-upload/model"
	"github.com/pengdst/golang-file-upload/service"
	"net/http"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func (a *AuthControllerImpl) Login(ctx *gin.Context) {
	var payload model.LoginPayload
	ctx.ShouldBind(&payload)

	ctx.JSON(http.StatusOK, model.WebResponse{
		Code:    http.StatusOK,
		Message: "success logged in",
		Data:    a.AuthService.Login(ctx, payload),
	})
}

func (a *AuthControllerImpl) RefreshToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.WebResponse{
		Code:    http.StatusOK,
		Message: "success refresh access token",
		Data:    a.AuthService.RefreshToken(ctx),
	})
}

func (a *AuthControllerImpl) Register(ctx *gin.Context) {
	var payload model.RegisterPayload
	ctx.ShouldBind(&payload)

	a.AuthService.Register(ctx, payload)

	ctx.JSON(http.StatusOK, model.WebResponse{
		Code:    http.StatusOK,
		Message: "success register new user",
	})
}

func (a *AuthControllerImpl) Logout(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{AuthService: authService}
}
