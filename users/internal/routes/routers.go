package routers

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine {
	g := gin.Default()
	UserRouters(g)
	return g
}
