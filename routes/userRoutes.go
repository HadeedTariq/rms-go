package routes

import (
	"rms-platform/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/auth")
	{
		userGroup.POST("/create", controllers.CreateUser)
	}
}
