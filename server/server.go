package server

import (
	"log"

	"github.com/Eduardo-Lisboa/api-go-gin/config"
	"github.com/Eduardo-Lisboa/api-go-gin/server/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() *Server {
	return &Server{
		port:   config.PortServer,
		server: gin.Default(),
	}
}

func (s *Server) Start() {

	router := routes.ConfigureRoutes(s.server)
	router.Use(cors.Default())
	log.Print("Server is running on port: ", s.port)

	log.Fatal(router.Run(":" + s.port))

}
