package handlers

import (
	"fmt"

	"github.com/MousaZa/e-vet/models"
	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	var p models.Product
	err := c.ShouldBind(p)
	if err != nil {
		fmt.Printf("Error parsing request body, %s", err)
	}

}
