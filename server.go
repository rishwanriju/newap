package main

import (
	"newsapp/controller"
	"newsapp/db"
	"newsapp/service"

	"github.com/gin-gonic/gin"
)

var (
	newsRepository db.NewsRepository         = db.NewNewsRepository()
	newsService    service.NewsService       = service.New(newsRepository)
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

	server.PUT("/news/:id", func(ctx *gin.Context) {

		ctx.JSON(200, newsController.Update(ctx))
	})

	server.DELETE("/news/:id", func(ctx *gin.Context) {

		ctx.JSON(200, newsController.Delete(ctx))
	})

	server.Run(":3000")
}
