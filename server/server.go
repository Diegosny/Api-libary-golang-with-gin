package server

import (
	"api/server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer () Server {
	return Server{
		port: ":8000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Fatal(router.Run(s.port))
}