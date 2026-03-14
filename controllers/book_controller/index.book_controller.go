package book_controller

import (
	"net/http"
	"restapi-gin/database"
	"restapi-gin/responses"

	"github.com/gin-gonic/gin"
)

func GetAllBook(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "book",
	})
}
func GetBookById(ctx *gin.Context) {
	id := ctx.Param("id")
	book := new(responses.BookResponse)
	errDB := database.DB.Table("books").Where("id = ?", id).Find(&book).Error

	if book.ID == nil {
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
		"data":    book,
	})
}
