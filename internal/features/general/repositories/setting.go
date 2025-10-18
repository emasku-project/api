package repositories

import (
	"errors"
	"strconv"

	"api/internal/features/general/models"
	"gorm.io/gorm"
)

type Setting struct {
	db *gorm.DB
}

func NewSetting(db *gorm.DB) *Setting {
	return &Setting{db: db}
}

func (s *Setting) GetTaxByUserId(userId uint) (float64, error) {
	var taxStr string
	if err := s.db.Model(&models.Setting{}).Where(
		"user_id = ? and key = ?",
		userId, "tax",
	).Select("value").First(&taxStr).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			taxStr = "5.0"
		} else {
			return 0.0, err
		}
	}

	if tax, err := strconv.ParseFloat(taxStr, 64); err != nil {
		return 0.0, err
	} else {
		return tax, nil
	}
}
