package entity

import (
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Title      string  `json:"title"`
	Total      float64 `json:"total"`
	Attachment string  `json:"attachment"`
}
