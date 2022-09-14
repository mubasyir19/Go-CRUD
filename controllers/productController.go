package controllers

import (
	"go-crud/initializers"
	"go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductCreate(ctx *gin.Context) {
	var body struct {
		Name  string
		Desc  string
		Price int64
	}
	ctx.Bind(&body)

	product := models.Product{
		Name:  body.Name,
		Desc:  body.Desc,
		Price: body.Price,
	}
	result := initializers.DB.Create(&product)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": product,
	})
}

func Products(ctx *gin.Context) {
	var products []models.Product

	initializers.DB.Find(&products)

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
