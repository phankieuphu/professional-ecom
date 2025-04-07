package services

import "github.com/gin-gonic/gin"

func InitOrder(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Init order success",
	})
}
