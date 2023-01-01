package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Role     int
	Password string
}
