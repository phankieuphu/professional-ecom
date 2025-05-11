package routers

import "github.com/gin-gonic/gin"

func PaymentRouters(r *gin.Engine) {
	group := r.Group("/payment")
	{
		group.GET("/list-payment", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"paypal": "https://paypal.com",
			})
		})

	}
}
