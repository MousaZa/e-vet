package main

import (
	"github.com/MousaZa/e-vet/server"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	s := server.NewServer(r)
	s.RunServer()
}
