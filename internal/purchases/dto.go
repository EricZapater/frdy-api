package purchases

import "time"

// PurchaseHeaderRequest represents the request payload for creating/updating a purchase header
type PurchaseHeaderRequest struct {
	Code      string    `json:"code" example:"PH-001" description:"Purchase header code"`
	SupplierName string    `json:"supplier_name" binding:"required" example:"Supplier A" description:"Name of the supplier"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z" description:"Purchase creation date"`
	Received bool `json:"received"`
}

// PurchaseDetailRequest represents the request payload for creating/updating a purchase detail
type PurchaseDetailRequest struct {
	PurchaseHeaderID string  `json:"purchase_header_id" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000" description:"Purchase header ID (UUID)"`
	ItemID          string  `json:"item_id" binding:"required" example:"123e4567-e89b-12d3-a456-426614174001" description:"Item ID (UUID)"`
	Quantity        int     `json:"quantity" binding:"required,min=1" example:"10" description:"Quantity of items"`
	Cost            float64 `json:"cost" binding:"required,min=0" example:"15.50" description:"Cost per unit"`
	Amount          float64 `json:"amount" binding:"required,min=0" example:"155.00" description:"Total amount (quantity * cost)"`
}