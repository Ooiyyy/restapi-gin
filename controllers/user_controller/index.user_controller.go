package user_controller

import (
	"restapi-gin/database"
	"restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {
	// isValidated := true
	// if !isValidated {

	// 	ctx.AbortWithStatusJSON(400, gin.H{
	// 		"message": "bad request, some field not validated",
	// 	})
	// 	return
	// }

	users := new([]models.User)

	err := database.DB.Table("users").Find(&users).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": users,
	})
}
