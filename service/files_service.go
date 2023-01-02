package service

import (
	"errors"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/pengdst/golang-file-upload/config"
	"github.com/pengdst/golang-file-upload/model"
	"github.com/pengdst/golang-file-upload/model/entity"
	"github.com/pengdst/golang-file-upload/repository"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FilesService interface {
	Upload(c *gin.Context, file *multipart.FileHeader) model.File
}

type FilesServiceImpl struct {
	env            *config.EnvConf
	fileUploadRepo repository.FileUploadRepository
}

func NewFilesService(env *config.EnvConf, fileUploadRepo repository.FileUploadRepository) FilesService {
	return &FilesServiceImpl{
		env:            env,
		fileUploadRepo: fileUploadRepo,
	}
}

func (f *FilesServiceImpl) Upload(c *gin.Context, file *multipart.FileHeader) model.File {
	uploadPath := "public/uploads"
	_, err := os.Stat(uploadPath)
	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	filePath := filepath.Join(uploadPath, file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		panic(err)
	}

	detectFile, err := mimetype.DetectFile(filePath)
	if err != nil {
		panic(err)
	}

	fileUpload := f.fileUploadRepo.Save(entity.FileUpload{
		Name:          file.Filename,
		MimeType:      detectFile.String(),
		FullPath:      filePath,
		RootDirectory: uploadPath,
	})

	return model.File{
		FileName:  fileUpload.Name,
		Directory: fileUpload.RootDirectory,
		MimeType:  fileUpload.MimeType,
		FileUrl:   fmt.Sprintf("%s/%s", f.env.BaseUrl, fileUpload.FullPath),
	}
}
