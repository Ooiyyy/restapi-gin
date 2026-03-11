package user_controller

import (
	"net/http"
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

func GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)
	errDB := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if errDB != nil || user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "data transmitted",
		"data":    user,
	})
}

func Store(ctx *gin.Context) {

}

func UpdateById(ctx *gin.Context) {

}

func DeleteById(ctx *gin.Context) {

}
