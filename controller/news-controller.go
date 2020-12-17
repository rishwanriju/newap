package controller

import (
	"newsapp/entity"
	"newsapp/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type NewsController interface {
	FindAll() []entity.News
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controller struct {
	service service.NewsService
}

var validate *validator.Validate

func New(service service.NewsService) NewsController {
	validate = validator.New()
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.News {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var news entity.News
	err := ctx.ShouldBindJSON(&news)
	if err != nil {
		return err
	}
	err = validate.Struct(news)
	if err != nil {
		return err
	}
	c.service.Save(news)
	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
	var news entity.News
	err := ctx.ShouldBindJSON(&news)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	news.ID = id

	err = validate.Struct(news)
	if err != nil {
		return err
	}
	c.service.Update(news)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var news entity.News
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	news.ID = id
	c.service.Delete(news)
	return nil
}
