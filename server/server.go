package server

import (
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

type Server struct {
	R  *gin.Engine
	DB *firestore.Client
}

func NewServer(r *gin.Engine, db *firestore.Client) *Server {
	return &Server{R: r, DB: db}
}

func (s *Server) SetupRoutes() {
	sg := s.R.Group("/stock")
	pg := sg.Group("/product")
	pg.POST("/")
}

func (s *Server) RunServer() {
	s.R.Run(":8080")
}
