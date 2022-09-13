package server

import (
	"github.com/gin-gonic/gin"
)

//TODO implement high-end server logic
//will need to add user stuff eventually
type Server struct {
	router *gin.Engine
	wh     WodHandler
}

func NewServer(router *gin.Engine, wh WodHandler) *Server {
	return &Server{
		router: router,
		wh:     wh,
	}
}

func (s *Server) handleIndex(c *gin.Context) {
	//add handle index here
	//may be a better way if we're just serving static files, though
}

func (s *Server) Run() {
	//register index
	s.router.GET("/", s.handleIndex)

	//register routes
	s.registerWodRoutes()

	s.router.Run(":8080")
}
