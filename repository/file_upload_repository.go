package repository

import (
	"github.com/pengdst/golang-file-upload/model/entity"
	"gorm.io/gorm"
)

type FileUploadRepository interface {
	Save(upload entity.FileUpload) entity.FileUpload
}

type FileUploadRepositoryImpl struct {
	DB *gorm.DB
}

func NewFileUploadRepository(DB *gorm.DB) FileUploadRepository {
	return &FileUploadRepositoryImpl{DB: DB}
}

func (f *FileUploadRepositoryImpl) Save(upload entity.FileUpload) entity.FileUpload {
	f.DB.Create(&upload)

	return upload
}
