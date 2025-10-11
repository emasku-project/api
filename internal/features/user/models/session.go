package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model

	UserId uint   `json:"user_id"`
	User   User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Token  string `json:"token"`
}
