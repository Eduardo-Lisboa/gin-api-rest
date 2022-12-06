package controllers

import (
	"net/http"

	"github.com/Eduardo-Lisboa/api-go-gin/database"
	"github.com/Eduardo-Lisboa/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func SearchForUser(c *gin.Context) {
	var user models.User
	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Username: input.Username, Password: input.Password}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UserLogin(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if user.Password != input.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is incorrect!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user,
		"message": "autenticated",
	})

}
