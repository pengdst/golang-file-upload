package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pengdst/golang-file-upload/model"
	"github.com/pengdst/golang-file-upload/service"
	"net/http"
)

type FilesController interface {
	Upload(c *gin.Context)
}

type FilesControllerImpl struct {
	fileService service.FilesService
}

func (f *FilesControllerImpl) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:    http.StatusOK,
		Message: "Success upload file",
		Data:    f.fileService.Upload(c, file),
	})
}

func NewFilesController(fileService service.FilesService) FilesController {
	return &FilesControllerImpl{
		fileService: fileService,
	}
}
