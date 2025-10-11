package models

import (
	"time"

	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model

	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Symbol    string    `json:"symbol"`
	UpdatedAt time.Time `json:"updated_at"`
}
