package web

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/blocks"
	"github.com/pengdst/golang-file-upload/repository"
	"net/http"
)

type AuthController interface {
	Login(ctx *gin.Context)
	LoginProcess(ctx *gin.Context)
	Register(ctx *gin.Context)
	RegisterProcess(ctx *gin.Context)
}

type AuthControllerImpl struct {
	views          *blocks.Blocks
	UserRepository repository.UserRepository
}

func (a *AuthControllerImpl) Login(ctx *gin.Context) {
	data := map[string]interface{}{
		"Title":    "Login",
		"ImageUrl": "http://www.w3.org/2000/svg",
	}
	err := a.views.ExecuteTemplate(ctx.Writer, "auth/login", "guest", data)
	if err != nil {
		panic(err)
	}
}

func (a *AuthControllerImpl) LoginProcess(ctx *gin.Context) {
	/*if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}*/

	ctx.Redirect(http.StatusFound, "/")
}

func (a *AuthControllerImpl) Register(ctx *gin.Context) {
	data := map[string]interface{}{
		"Title":    "Login",
		"ImageUrl": "http://www.w3.org/2000/svg",
	}
	err := a.views.ExecuteTemplate(ctx.Writer, "auth/register", "guest", data)
	if err != nil {
		panic(err)
	}
}

func (a *AuthControllerImpl) RegisterProcess(ctx *gin.Context) {
	/*	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/register")
		return
	}*/

	ctx.Redirect(http.StatusFound, "/login")
}

func NewAuthController(views *blocks.Blocks, userRepo repository.UserRepository) AuthController {
	return &AuthControllerImpl{
		views:          views,
		UserRepository: userRepo,
	}
}
