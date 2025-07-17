package items

import "github.com/google/uuid"

type Item struct {
	ID 			uuid.UUID `json:"id" binding:"required"`
	Code        string  `json:"code" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Cost        float64 `json:"cost" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	IsActive	bool    `json:"is_active" binding:"required"`
}