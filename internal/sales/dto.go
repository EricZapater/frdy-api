package sales

type SalesHeaderRequest struct {
	Code         string `json:"code" binding:"required"`
	CustomerName string `json:"customer_name" binding:"required"`
}

type SalesDetailRequest struct {
	SalesHeaderID string  `json:"sales_header_id" binding:"required"`
	ItemID        string  `json:"item_id" binding:"required"`
	Quantity      int     `json:"quantity" binding:"required"`
	Price         float64 `json:"price" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
}