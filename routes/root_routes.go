package routes

import (
	"github.com/gin-gonic/gin"
	"session-root/controllers"
)

func RootRoutes(r *gin.Engine) {
	rootGroup := r.Group("/")
	{
		rootGroup.GET("health/", controllers.CheckHealth())
	}
}
