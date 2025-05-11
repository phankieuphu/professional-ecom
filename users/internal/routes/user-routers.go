package routers

import (
	services "github.com/phankieuphu/ecom-user/internal/services"

	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.POST("/register-user", services.RegisterUser)
}
