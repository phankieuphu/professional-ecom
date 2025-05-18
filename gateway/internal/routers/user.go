package routers

import (
	"gateway/internal/services"

	"github.com/gin-gonic/gin"
)

func UserRegister(r *gin.Engine) {
	group := r.Group("/user")
	{
		group.GET("/info/:username", services.GetUser)
		group.POST("/register", services.RegisterUser)
		group.PATCH("/update/:username", services.UpdateUser)
	}

}
