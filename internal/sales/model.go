package sales

import "github.com/google/uuid"

type SalesHeader struct {
	ID           uuid.UUID `json:"id" binding:"required"`
	Code         string    `json:"code" binding:"required"`
	CustomerName string    `json:"customer_name" binding:"required"`
	CustomerPhone string `json:"customer_phone"`
	CreatedAt    string    `json:"created_at" binding:"required"`
	Sent         bool      `json:"sent" binding:"required"`
}

type SalesDetail struct {
	ID              uuid.UUID  `json:"id" binding:"required"`
	SalesHeaderID   string  `json:"sales_header_id" binding:"required"`
	ItemID          string  `json:"item_id" binding:"required"`
	ItemCode        string  `json:"item_code" binding:"required"`
	ItemDescription string  `json:"item_description" binding:"required"`
	Quantity        int     `json:"quantity" binding:"required"`
	Price           float64 `json:"price" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
}