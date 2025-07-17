package items

type ItemRequest struct {
	Code        string  `json:"code" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Cost        float64 `json:"cost" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}