package user_controller

import (
	"errors"
	"net/http"
	"restapi-gin/database"
	"restapi-gin/models"
	"restapi-gin/requests"
	"restapi-gin/responses"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

// func Store(ctx *gin.Context) {
// 	userReq := new(requests.UserRequest)
// 	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
// 		ctx.JSON(400, gin.H{
// 			"error":   "isian form kurang lengkap",
// 			"errors":  errReq.Error(),
// 			"message": "data yang diberikan tidak valid",
// 		})
// 		return
// 	}
// 	// Manual validation for unique email
// 	userEmailExist := new(models.User)
// 	database.DB.Table("users").Where("email = ?", userReq.Email).First(&userEmailExist)
// 	if userEmailExist.Email != nil {
// 		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
// 			"message": "email already used",
// 		})
// 		return
// 	}

// 	user := new(models.User)
// 	user.Name = &userReq.Name
// 	user.Email = &userReq.Email
// 	user.Address = &userReq.Address
// 	user.BornDate = &userReq.BornDate

//		errDb := database.DB.Table("users").Create(&user).Error
//		if errDb != nil {
//			ctx.JSON(500, gin.H{
//				"message": "can't create data",
//			})
//			return
//		}
//		ctx.JSON(200, gin.H{
//			"message": "data created",
//			"data":    user,
//		})
//	}
func Store(ctx *gin.Context) {

	var userReq requests.UserRequest

	if err := ctx.ShouldBind(&userReq); err != nil {

		var ve validator.ValidationErrors
		fieldErrors := make(map[string]string)

		if errors.As(err, &ve) {
			for _, fe := range ve {

				field := strings.ToLower(fe.Field())

				switch fe.Tag() {

				case "required":
					fieldErrors[field] = field + " wajib diisi"

				case "email":
					fieldErrors[field] = "format email tidak valid"

				case "min":
					fieldErrors[field] = field + " minimal " + fe.Param() + " karakter"

				default:
					fieldErrors[field] = "nilai tidak valid"
				}
			}
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Data yang diberikan tidak valid",
			"errors":  fieldErrors,
		})
		return
	}

	// cek email unique
	var userExist models.User
	result := database.DB.Where("email = ?", userReq.Email).First(&userExist)

	if result.RowsAffected > 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "validation error",
			"errors": gin.H{
				"email": "email already used",
			},
		})
		return
	}

	user := models.User{
		Name:     &userReq.Name,
		Email:    &userReq.Email,
		Address:  &userReq.Address,
		BornDate: &userReq.BornDate,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't create data",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "data created",
		"data":    user,
	})
}

func UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)
	userReq := new(requests.UserRequest)
	userEmailExist := new(models.User)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "data not found",
		})
		return
	}
	// email exists
	errUserEmailExist := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&userEmailExist).Error
	if errUserEmailExist != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if userEmailExist.Email != nil && *user.ID != *userEmailExist.ID {
		ctx.JSON(400, gin.H{
			"message": "email already used",
		})
		return
	}

	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate
	errUpdate := database.DB.Table("users").Where("id = ?", id).Updates(&user).Error
	if errUpdate != nil {
		ctx.JSON(500, gin.H{
			"message": "Can't update data",
		})
		return
	}

	//custom response field
	userResponse := responses.UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}
	ctx.JSON(200, gin.H{
		"message": "data updated",
		"data":    userResponse,
	})
}

func DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	//validation user id
	user := new(models.User)
	errFind := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errFind != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if user.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "data not found",
		})
		return
	}

	errDb := database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(&models.User{}).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   errDb.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "data deleted",
	})
}
