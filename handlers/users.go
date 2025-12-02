package handlers

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/MousaZa/e-vet/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserWithDB(db *firestore.Client) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var u models.User
		err := ctx.ShouldBindJSON(&u)
		if err != nil {
			fmt.Printf("Error fetching user data: %s\n", err)
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		hp, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("Error hashing passowrd: %s\n", err)
			ctx.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		u.Password = string(hp)
		_, _, err = db.Collection("Users").Add(ctx.Request.Context(), u)
		if err != nil {
			fmt.Printf("Error storing the user: %s\n", err)
			ctx.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		ctx.Writer.WriteHeader(http.StatusOK)
	}
}

func LoginWithDB(db *firestore.Client) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var lr models.LoginRequest
		ctx.ShouldBindJSON(&lr)
		q := db.Collection("Users").Where("Email", "==", lr.Email).Documents(ctx.Request.Context())
		docs, err := q.GetAll()
		if err != nil {
			fmt.Printf("Something went wrong: %s\n", err)
			ctx.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		doc := docs[0]
		var u models.User

		err = doc.DataTo(&u)
		if err != nil {
			fmt.Printf("Something went wrong: %s\n", err)
			ctx.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(lr.Password))
		if err != nil {
			fmt.Printf("Wrong Password: %s\n", err)
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx.JSON(http.StatusOK, models.LoginResponse{Username: u.Username, Token: ""})
	}
}
