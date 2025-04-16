package services

import (
	"fmt"
	"gateway/database"

	"github.com/gin-gonic/gin"
)

func InitOrder(c *gin.Context) {
	dbInstance := database.GetDb()
	user := dbInstance.Table("user").Select("")
	fmt.Println(user)
	c.JSON(200, gin.H{
		"Message": "Init order success",
	})
}
