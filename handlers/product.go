package handlers

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/MousaZa/e-vet/models"
	"github.com/gin-gonic/gin"
)

func AddProductWithDB(db *firestore.Client) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var p models.Product
		err := ctx.ShouldBindJSON(&p)
		if err != nil {
			fmt.Printf("Error parsing request body, %s", err)
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		p.IsActive = true
		// var u models.User
		// p.Register(&u)
		st := db.Collection("Stock")
		_, err1 := st.NewDoc().Create(ctx.Request.Context(), p)
		if err1 != nil {
			ctx.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		ctx.Writer.WriteHeader(http.StatusOK)
	}
}

func GetProductsWithDB(db *firestore.Client) func(*gin.Context) {
	return func(ctx *gin.Context) {
		st := db.Collection("Stock")
		docs := st.Documents(ctx.Request.Context())

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
				ctx.Writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = doc.DataTo(&r)
			if err != nil {
				fmt.Printf("Error getting the docs: %s", err)
				ctx.Writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			r.ID = doc.Ref.ID
			resp = append(resp, r)
		}
		fmt.Println(resp)
		ctx.JSON(http.StatusOK, resp)
	}
}

func DeleteProductWithDB(db *firestore.Client) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		doc := db.Collection("Stock").Doc(id)

		_, err := doc.Delete(ctx.Request.Context())
		if err != nil {
			fmt.Printf("Error deleting the product with id %s: %s", id, err)
			ctx.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		ctx.Writer.WriteHeader(http.StatusOK)
	}
}

type ConsumeRequest struct {
	ConsumeAmt int `json:"consume_amt"`
}

func ConsumeProductWithDB(db *firestore.Client) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var ca ConsumeRequest
		ctx.ShouldBindJSON(&ca)

		id := ctx.Param("id")
		doc := db.Collection("Stock").Doc(id)

		ds, err := doc.Get(ctx.Request.Context())
		if err != nil {
			fmt.Printf("Error getting the documnet: %s", err)
			ctx.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		var d models.Product
		ds.DataTo(&d)
		d.Quantity -= ca.ConsumeAmt

		_, err = doc.Set(ctx.Request.Context(), d)

		if err != nil {
			fmt.Printf("Error updating quantity: %s", err)
			ctx.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		ctx.Writer.WriteHeader(http.StatusOK)
	}
}
