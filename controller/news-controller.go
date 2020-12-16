package controller

import (
	"newsapp/entity"
	"newsapp/service"

	"github.com/gin-gonic/gin"
)

type NewsController interface {
	FindAll() []entity.News
	Save(ctx *gin.Context) entity.News
}

type controller struct {
	service service.NewsService
}

func New(service service.NewsService) NewsController {

	return &controller{

		service: service,
	}
}

func (c *controller) FindAll() []entity.News {

	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) entity.News {

	var newss entity.News
	ctx.BindJSON(&newss)
	c.service.Save(newss)
	return newss
}
