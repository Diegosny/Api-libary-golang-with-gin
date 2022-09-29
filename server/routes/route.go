package routes

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controller.ShowBook)
			books.POST("/", controller.CreateBook)
			books.GET("/", controller.ShowBooks)
			books.PUT("/:id", controller.UpdateBook)
			books.DELETE("/:id", controller.DeleteBook)
		}
	}

	return router
}
