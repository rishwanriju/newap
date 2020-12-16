package main

import (
	"newsapp/controller"
	"newsapp/service"

	"github.com/gin-gonic/gin"
)

var (
	newsService    service.NewsService       = service.New()
	newsController controller.NewsController = controller.New(newsService)
)

func main() {

	server := gin.Default()

	server.GET("/news", func(ctx *gin.Context) {

		ctx.JSON(200, newsController.FindAll())
	})

	server.POST("/news", func(ctx *gin.Context) {

		ctx.JSON(200, newsController.Save(ctx))
	})
	server.Run(":3000")
}
