package models

import "time"

type Staff struct {
    ID         string    `json:"id" gorm:"primaryKey"`
    Name       string    `json:"name"`
    Role       string    `json:"role"`
    Shift      string    `json:"shift"`
    EmployedOn time.Time `json:"employed_on"`
    UpdatedBy  string    `json:"updated_by"`
}