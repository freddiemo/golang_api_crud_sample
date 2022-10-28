package service

import "github.com/golang_api_crud_sample/entity"

type UserService interface {
	Save(entity.User) entity.User
}

type userService struct {
	users []entity.User
}

func New() UserService {
	return &userService{}
}

func (service *userService) Save(user entity.User) entity.User {
	service.users = append(service.users, user)
	return user
}
