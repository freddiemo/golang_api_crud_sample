package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/golang_api_crud_sample/entity"
	"github.com/golang_api_crud_sample/service"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context) (entity.User, error)
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.User {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) (entity.User, error) {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return entity.User{}, err
	}
	c.service.Save(user)
	return user, nil
}