package models

import "time"


type Staff struct {
    ID         string    `gorm:"primaryKey" json:"id"`
    Name       string    `json:"name"`
    Role       string    `json:"role"`
    Shift      string    `json:"shift"`
    UpdatedBy  string    `json:"updated_by"` // Who added this staff?
    CreatedAt  time.Time `gorm:"autoCreateTime" json:"employed_on"`
}