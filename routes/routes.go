package routes

import (
	"github.com/Eduardo-Lisboa/api-go-gin/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/produtos", controllers.Products)
	r.GET("/produtos/:id", controllers.SearchForProduct)
	r.POST("/produtos", controllers.CreatProduct)
	r.DELETE("/produtos/:id", controllers.DeleteProduct)
	r.PATCH("/produtos/:id", controllers.UpdateProduct)
	r.Use(cors.Default())
	r.Run(":8080")
}
