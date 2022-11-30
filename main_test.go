package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Eduardo-Lisboa/api-go-gin/controllers"
	"github.com/Eduardo-Lisboa/api-go-gin/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTest() *gin.Engine {
	r := gin.Default()
	return r

}

func TestStatusCode(t *testing.T) {
	database.ConectDatabase()
	r := SetupTest()
	r.GET("/produtos", controllers.Products)
	req, _ := http.NewRequest("GET", "/produtos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, 200, response.Code)

}
