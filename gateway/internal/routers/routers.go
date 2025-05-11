package routers

import "github.com/gin-gonic/gin"

func SetupRouters() *gin.Engine {
	router := gin.Default()
	UserRegister(router)
	PaymentRouters(router)
	OrderRouters(router)
	return router
}
