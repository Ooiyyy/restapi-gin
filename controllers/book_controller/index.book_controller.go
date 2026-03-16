package book_controller

import (
	"net/http"
	"restapi-gin/database"
	"restapi-gin/models"
	"restapi-gin/requests"
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
func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	book := new(models.Book)
	bookReq := new(requests.BookRequest)

	if errReq := ctx.ShouldBind(&bookReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}
	errDb := database.DB.Table("books").Where("id = ?", id).Find(&book).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server error",
		})
		return
	}

	if book.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "Data not found",
		})
		return
	}

	book.Title = &bookReq.Title
	book.Author = &bookReq.Author
	book.Description = &bookReq.Description
	book.Price = &bookReq.Price
	errUpdate := database.DB.Table("books").Where("id = ?", id).Updates(&book).Error
	if errUpdate != nil {
		ctx.JSON(500, gin.H{
			"message": "Can't update data",
		})
		return
	}
	//custom response field
	BookResponse := responses.BookResponse{
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		Price:       book.Price,
	}
	ctx.JSON(200, gin.H{
		"message": "data updated",
		"data":    BookResponse,
	})
}
