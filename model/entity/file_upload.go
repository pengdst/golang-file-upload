package entity

import "gorm.io/gorm"

type FileUpload struct {
	gorm.Model
	Name          string
	MimeType      string
	FullPath      string
	RootDirectory string
}
