package routers

import (
	"github.com/baguspanji/go-crud/controllers"
	"github.com/baguspanji/go-crud/middleware"
	"github.com/gin-gonic/gin"
)

func PostRouter(r *gin.Engine) {
	postGroup := r.Group("/post", middleware.RequireAuth)
	{
		postGroup.GET("/", controllers.GetPosts)
		postGroup.GET("/:id", controllers.GetPost)
		postGroup.POST("/", controllers.CreatePost)
		postGroup.PUT("/:id", controllers.UpdatePost)
		postGroup.DELETE("/:id", controllers.DeletePost)
	}
}
