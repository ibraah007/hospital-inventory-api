package models

// Ensure this says 'Item', not 'Inventory'
type Item struct {
	ID         string `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Quantity   int    `json:"quantity"`
	ExpiryDate string `json:"expiry_date"`
}