package models

import "time"

type Inventory struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Category   string    `json:"category"`   
	Department string    `json:"department" json:"department"` 
	Quantity   int       `json:"quantity"`
	ExpiryDate string    `json:"expiry_date" json:"expiry_date"` // YYYY-MM-DD
	CreatedAt  time.Time `json:"created_at"`
}