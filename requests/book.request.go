package requests

type BookRequest struct {
	Title       string  `json:"title"form:"title" binding:"required"`
	Author      string  `json:"author"form:"author" binding:"required"`
	Description string  `json:"description"form:"description" binding:"required"`
	Price       float32 `json:"price"form:"price" binding:"required"`
}
