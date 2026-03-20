package models

import "time"

type Item struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}