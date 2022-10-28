package service

import (
	"github.com/golang_api_crud_sample/entity"
	"github.com/golang_api_crud_sample/repository"
)

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
	Update(user entity.User)
	Delete(user entity.User)
}

type userService struct {
	userRepository repository.UserRepository
}

func New(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Save(user entity.User) entity.User {
	service.userRepository.Save(user)
	return user
}

func (service *userService) FindAll() []entity.User {
	return service.userRepository.FindAll()
}

func (service *userService) Update(user entity.User) {
	service.userRepository.Update(user)
}

func (service *userService) Delete(user entity.User) {
	service.userRepository.Delete(user)
}