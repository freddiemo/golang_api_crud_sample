package controller

import (
	"github.com/gin-gonic/gin"

	"strconv"

	"github.com/golang_api_crud_sample/entity"
	"github.com/golang_api_crud_sample/service"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context) (entity.User, error)
	Update(ctx *gin.Context) (entity.User, error)
	Delete(ctx *gin.Context) error
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

func (c *controller) Update(ctx *gin.Context) (entity.User, error) {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return entity.User{}, err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	user.ID = id
	c.service.Update(user)
	return user, nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var user entity.User
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	user.ID = id

	c.service.Delete(user)

	return nil
}