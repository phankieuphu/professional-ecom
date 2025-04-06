package routers

import "github.com/gin-gonic/gin"

func UserRegister(r *gin.Engine) {
	group := r.Group("/user")
	{
		group.GET("/info", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code": 0,
			})

		})

	}

}
