package purchases

import "time"

// PurchaseHeader represents a purchase header entity
type PurchaseHeader struct {
	ID        string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" description:"Purchase header ID (UUID)"`
	Code      string    `json:"code" example:"PH-001" description:"Purchase header code"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z" description:"Purchase creation date"`
}

// PurchaseDetail represents a purchase detail entity
type PurchaseDetail struct {
	ID              string  `json:"id" example:"123e4567-e89b-12d3-a456-426614174002" description:"Purchase detail ID (UUID)"`
	PurchaseHeaderID string  `json:"purchase_header_id" example:"123e4567-e89b-12d3-a456-426614174000" description:"Purchase header ID (UUID)"`
	ItemID          string  `json:"item_id" example:"123e4567-e89b-12d3-a456-426614174001" description:"Item ID (UUID)"`
	ItemCode        string  `json:"item_code" example:"ITEM-001" description:"Item code from joined table"`
	ItemDescription string  `json:"item_description" example:"Sample Item Description" description:"Item description from joined table"`
	Quantity        int     `json:"quantity" example:"10" description:"Quantity of items"`
	Cost            float64 `json:"cost" example:"15.50" description:"Cost per unit"`
	Amount          float64 `json:"amount" example:"155.00" description:"Total amount (quantity * cost)"`
}
