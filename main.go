package main

import (
	"github.com/baguspanji/go-crud/initializers"
	"github.com/baguspanji/go-crud/routers"
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

	routers.AuthRouter(r)
	routers.PostRouter(r)

	r.Run()
}
