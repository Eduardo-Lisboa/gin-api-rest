package main

import (
	"github.com/Eduardo-Lisboa/api-go-gin/database"
	"github.com/Eduardo-Lisboa/api-go-gin/routes"
)

func main() {
	database.ConectDatabase()

	routes.HandleRequests()

}
