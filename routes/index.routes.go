package routes

import (
	"restapi-gin/controllers/book_controller"
	"restapi-gin/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app
	//route user
	route.GET("/user", user_controller.GetAllUser)
	route.GET("/user/:id", user_controller.GetById)

	//route book
	route.GET("/book", book_controller.GetAllBook)

}
