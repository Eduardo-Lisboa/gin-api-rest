package routes

import (
	"github.com/Eduardo-Lisboa/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")
	{
		products := main.Group("products")
		{
			products.GET("/", controllers.Products)
			products.GET("/:id", controllers.SearchForProduct)
			products.POST("/", controllers.CreateProduct)
			products.DELETE("/:id", controllers.DeleteProduct)
			products.PATCH("/:id", controllers.UpdateProduct)
		}

	}
	return router

}
