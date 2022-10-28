package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/golang_api_crud_sample/entity"
	"github.com/golang_api_crud_sample/service"
)

type UserController interface {
	Save(ctx *gin.Context) entity.User
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	c.service.Save(user)
	return user
}