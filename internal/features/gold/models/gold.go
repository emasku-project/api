package models

import (
	"time"

	"api/internal/features/user/models"
	"gorm.io/gorm"
)

type Gold struct {
	gorm.Model

	UserId       uint        `json:"user_id"`
	User         models.User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PurchaseDate time.Time   `json:"purchase_date"`
	Weight       float64     `json:"weight"`
	Price        float64     `json:"price"`
	Note         string      `json:"note"`
}
