package repository

import (
	"context"
	"errors"
	"github.com/pengdst/golang-file-upload/model/entity"
	"github.com/pengdst/golang-file-upload/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) (*entity.User, error)
	Verify(ctx context.Context, email string, password string) (*entity.User, error)
	GetAll(ctx context.Context) []entity.User
	Get(ctx context.Context, userId int) (*entity.User, error)
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}

func (u *UserRepositoryImpl) Verify(ctx context.Context, email string, password string) (*entity.User, error) {
	var user entity.User
	result := u.Db.Where(&entity.User{Email: email}).First(&user)

	if result.Error != nil {
		return nil, errors.New("user unauthenticated")
	}

	err := utils.VerifyPassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepositoryImpl) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	result := u.Db.Create(&user).Error

	if result != nil {
		return nil, result
	}

	return &user, nil
}

func (u *UserRepositoryImpl) GetAll(ctx context.Context) []entity.User {
	var users []entity.User
	u.Db.WithContext(ctx).Find(&users)

	return users
}

func (u *UserRepositoryImpl) Get(ctx context.Context, userId int) (*entity.User, error) {
	var user entity.User
	result := u.Db.WithContext(ctx).First(&user, userId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
