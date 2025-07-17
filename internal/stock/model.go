package stock

type Stock struct {
	ID              string `json:"id" db:"id"`
	ItemID          string `json:"item_id" db:"item_id"`
	ItemCode        string `json:"item_code" db:"item_code"`
	ItemDescription string `json:"item_description" db:"item_description"`
	Quantity        int    `json:"quantity" db:"quantity"`
}