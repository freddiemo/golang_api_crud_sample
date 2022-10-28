package repository

import (
	"gorm.io/gorm"

	"github.com/golang_api_crud_sample/entity"
)

type UserRepository interface {
	Save(user entity.User)
	FindAll() []entity.User
}

type userRepo struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo {
		connection: db,
	}
}

func (userRepo *userRepo) Save(user entity.User) {
	userRepo.connection.Create(&user)
}

func (userRepo *userRepo) FindAll() []entity.User {
	var users []entity.User
	userRepo.connection.Find(&users)
	// db.connection.Set("gorm:auto_preload", true).Find(&Users)
	return users
}