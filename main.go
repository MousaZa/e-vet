package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/MousaZa/e-vet/server"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {
	r := gin.New()

	ctx := context.Background()
	opt := option.WithCredentialsFile("e-vet.json")
	client, err := firestore.NewClient(ctx, "e-vet-cd914", opt)
	if err != nil {
		panic(fmt.Sprintf("Error connecting with firestore: %s", err))
	}

	s := server.NewServer(r, client)
	s.RunServer()
}
