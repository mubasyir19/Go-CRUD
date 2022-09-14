package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(ctx *gin.Context) {
	// GET data from req body
	var body struct {
		Body  string
		Title string
	}

	ctx.Bind(&body)

	// Create a post
	// post := models.Post{Title: "Hello", Body: "Post body"}
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	// Return
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(ctx *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond
	ctx.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsShow(ctx *gin.Context) {
	// Get id off URL
	id := ctx.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond
	ctx.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(ctx *gin.Context) {
	// GET id off URL
	id := ctx.Param("id")

	//GET Data off req body
	var body struct {
		Title string
		Body  string
	}
	ctx.Bind(&body)

	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//Update
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//Respond
	ctx.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(ctx *gin.Context) {
	// Get id off URL
	id := ctx.Param("id")

	//Delete
	initializers.DB.Delete(&models.Post{}, id)

	//Respond
	ctx.JSON(200, gin.H{
		"message": "Success delete data",
	})
}
