package server

import "github.com/gin-gonic/gin"

type Server struct {
	R *gin.Engine
}

func (s *Server) SetupRoutes() {
	sg := s.R.Group("/stock")
	pg := sg.Group("/product")
	pg.POST("/")
}

func NewServer(r *gin.Engine) *Server {
	return &Server{R: r}
}

func (s *Server) RunServer() {
	s.R.Run(":8080")
}
