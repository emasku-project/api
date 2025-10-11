package models

import (
	"time"

	"gorm.io/gorm"
)

type Currency struct {
	gorm.Model

	Base string    `json:"base"`
	Date time.Time `json:"date"`
	Rate float64   `json:"rate"`
}
