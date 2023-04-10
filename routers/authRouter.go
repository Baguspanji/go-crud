package routers

import (
	"github.com/baguspanji/go-crud/controllers"
	"github.com/baguspanji/go-crud/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.GET("/", middleware.RequireAuth, controllers.AuthUser)
		authGroup.POST("/register", controllers.AuthRegister)
		authGroup.POST("/login", controllers.AuthLogin)
	}
}
