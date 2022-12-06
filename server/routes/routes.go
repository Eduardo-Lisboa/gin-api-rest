package routes

import (
	"github.com/Eduardo-Lisboa/api-go-gin/config"
	"github.com/Eduardo-Lisboa/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group(config.ApiName)
	{
		products := main.Group(config.Products)
		{
			products.GET("/", controllers.Products)
			products.GET("/:id", controllers.SearchForProduct)
			products.POST("/", controllers.CreateProduct)
			products.DELETE("/:id", controllers.DeleteProduct)
			products.PATCH("/:id", controllers.UpdateProduct)

		}
		users := main.Group(config.Users)
		{
			users.GET("/", controllers.GetUsers)
			users.GET("/:id", controllers.SearchForUser)
			users.POST("/", controllers.CreateUser)
			users.GET("/login", controllers.UserLogin)
		}

	}
	return router

}
