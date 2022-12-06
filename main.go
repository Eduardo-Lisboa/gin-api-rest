package main

import (
	"github.com/Eduardo-Lisboa/api-go-gin/database"
	"github.com/Eduardo-Lisboa/api-go-gin/server"
)

func main() {

	database.ConectDatabase()

	server := server.NewServer()

	server.Start()

}
