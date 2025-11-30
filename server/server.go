package server

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/MousaZa/e-vet/handlers"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type Server struct {
	R  *gin.Engine
	DB *firestore.Client
}

func New() *Server {
	r := gin.New()

	ctx := context.Background()
	opt := option.WithCredentialsFile("../e-vet.json")
	db, err := firestore.NewClient(ctx, "e-vet-cd914", opt)
	if err != nil {
		panic(fmt.Sprintf("Error connecting with firestore: %s", err))
	}
	s := &Server{R: r, DB: db}
	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.R.POST("/stock/products", handlers.AddProductWithDB(s.DB))
	s.R.GET("/stock/products", handlers.GetProductsWithDB(s.DB))
}

func (s *Server) Run() {
	s.R.Run(":8080")
}
