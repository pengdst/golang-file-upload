package entity

import "gorm.io/gorm"

type AccessToken struct {
	Role      string
	AuthToken string
	UserID    string
	User      User
	gorm.Model
}
