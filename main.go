package main

import (
	"github.com/baguspanji/go-crud/controllers"
	"github.com/baguspanji/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// wellcome route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my API",
		})
	})

	postGroup := r.Group("/post")
	{
		postGroup.GET("/", controllers.GetPosts)
		postGroup.GET("/:id", controllers.GetPost)
		postGroup.POST("/", controllers.CreatePost)
		postGroup.PUT("/:id", controllers.UpdatePost)
		postGroup.DELETE("/:id", controllers.DeletePost)
	}

	r.Run()
}
