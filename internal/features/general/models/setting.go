package models

import (
	"api/internal/features/user/models"
	"gorm.io/gorm"
)

type Setting struct {
	gorm.Model

	UserId uint        `json:"user_id"`
	User   models.User `json:"user" gorm:"onUpdate:CASCADE,onDelete:CASCADE;"`
	Key    string      `json:"key"`
	Value  string      `json:"value"`
}
