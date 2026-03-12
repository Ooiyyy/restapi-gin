package user_controller

import (
	"net/http"
	"restapi-gin/database"
	"restapi-gin/models"
	"restapi-gin/requests"
	"restapi-gin/responses"

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
	user := new(responses.UserResponse)
	errDB := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}
	if errDB != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "data transmitted",
		"data":    user,
	})
}

func Store(ctx *gin.Context) {
	userReq := new(requests.UserRequest)
	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}
	// Manual validation for unique email
	userEmailExist := new(models.User)
	database.DB.Table("users").Where("email = ?", userReq.Email).First(&userEmailExist)
	if userEmailExist.Email != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "email already used",
		})
		return
	}

	user := new(models.User)
	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	errDb := database.DB.Table("users").Create(&user).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "can't create data",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "data created",
		"data":    user,
	})
}

func UpdateById(ctx *gin.Context) {

}

func DeleteById(ctx *gin.Context) {

}
