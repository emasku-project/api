package repositories

import (
	"errors"
	"fmt"
	"strconv"

	"api/internal/features/general/domains"
	"api/internal/features/general/models"
	"gorm.io/gorm"
)

type Setting struct {
	db *gorm.DB
}

func NewSetting(db *gorm.DB) *Setting {
	return &Setting{db: db}
}

func (r *Setting) GetTaxByUserId(userId uint) (float64, error) {
	var taxStr string
	if err := r.db.Model(&models.Setting{}).Where(
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

func (r *Setting) UpdateTaxByUserId(userId uint, tax float64) (*domains.Setting, error) {
	var setting models.Setting
	if err := r.db.Where(
		models.Setting{
			UserId: userId,
			Key:    "tax",
		},
	).Assign(
		models.Setting{
			Value: fmt.Sprintf("%.2f", tax),
		},
	).FirstOrCreate(&setting).Error; err != nil {
		return nil, err
	} else {
		return domains.FromSettingModel(&setting), nil
	}
}
