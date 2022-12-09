package web

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/blocks"
	"github.com/pengdst/golang-file-upload/repository"
)

type HomeController interface {
	Index(ctx *gin.Context)
}

type HomeControllerImpl struct {
	views    *blocks.Blocks
	UserRepo repository.UserRepository
}

func (h *HomeControllerImpl) Index(ctx *gin.Context) {
	data := map[string]interface{}{
		"Users": h.UserRepo.GetAll(ctx),
	}

	err := h.views.ExecuteTemplate(ctx.Writer, "index", "admin", data)
	if err != nil {
		panic(err)
	}
}

func NewHomeController(views *blocks.Blocks, userRepo repository.UserRepository) HomeController {
	return &HomeControllerImpl{
		views:    views,
		UserRepo: userRepo,
	}
}
