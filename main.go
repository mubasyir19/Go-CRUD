package main

import (
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// Template HTML
	router.LoadHTMLGlob("templates/*")
	router.LoadHTMLFiles("templates/index.html")

	// Example Router
	router.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(200,"index.html", gin.H{"name":"Mahdy Mubasyir","age":20})
	})

	// Router Post
	router.POST("/posts", controllers.PostsCreate)
	router.PUT("/posts/:id", controllers.PostsUpdate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)
	router.DELETE("/posts/:id", controllers.PostsDelete)

	// Router Product
	router.POST("/product", controllers.ProductCreate)
	router.GET("/product", controllers.Products)

	router.Run()
}
