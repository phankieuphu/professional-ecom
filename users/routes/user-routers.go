package routers

import (
	"users/services"

	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.POST("/register-user", services.RegisterUser)
	userGroup.GET("/get-profile/:username", services.GetUserProfile)
}
