package handlers

import (
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
		st := db.Collection("Stock")
		_, err1 := st.NewDoc().Create(c.Request.Context(), p)
		if err1 != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
	}
}
