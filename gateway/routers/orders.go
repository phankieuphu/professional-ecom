package routers

import (
	"gateway/services"

	"github.com/gin-gonic/gin"
)

func OrderRouters(r *gin.Engine) {
	group := r.Group("/orders")
	{
		// validate voucher here
		group.POST("/init-order", services.InitOrder)
		group.POST("/confirm-order", func(c *gin.Context) {
			c.JSON(201,
				gin.H{
					"Message": "Confirm order success",
				})
		})
		group.GET("/get-status", func(c *gin.Context) {
			c.JSON(200,
				gin.H{
					"status": "Pending",
				},
			)
		})
	}
}
