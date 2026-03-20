package models

import "time"

type Staff struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Shift     string    `json:"shift"`
	CreatedAt time.Time `json:"created_at"` // Automatically set by the DB
}