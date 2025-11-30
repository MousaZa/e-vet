package handlers

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/MousaZa/e-vet/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddProductWithDB(db *firestore.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		var p models.Product
		err := c.ShouldBindJSON(&p)
		if err != nil {
			fmt.Printf("Error parsing request body, %s", err)
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		p.ID = int64(uuid.New()[0])
		p.IsActive = true
		st := db.Collection("Stock")
		_, err1 := st.NewDoc().Create(c.Request.Context(), p)
		if err1 != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
	}
}

func GetProductsWithDB(db *firestore.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		st := db.Collection("Stock")
		docs := st.Documents(c.Request.Context())

		var resp []models.Product
		for {
			n, err := docs.Next()
			if err != nil {
				break
			}
			var r models.Product
			doc, err := n.Ref.Get(context.TODO())
			if err != nil {
				fmt.Printf("Error getting the docs: %s", err)
				c.Writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = doc.DataTo(&r)
			if err != nil {
				fmt.Printf("Error getting the docs: %s", err)
				c.Writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			resp = append(resp, r)
		}
		fmt.Println(resp)
		c.JSON(http.StatusOK, resp)
	}
}
