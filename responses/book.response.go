package responses

type BookResponse struct {
	ID          *int     `json:"id"`
	Title       *string  `json:"title"`
	Author      *string  `json:"author"`
	Description *string  `json:"description"`
	Price       *float32 `json:"price"`
}
