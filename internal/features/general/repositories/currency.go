package repositories

import (
	"time"

	"api/internal/features/general/domains"
	"api/internal/features/general/models"
	"gorm.io/gorm"
)

type Currency struct {
	db *gorm.DB
}

func NewCurrency(
	db *gorm.DB,
) *Currency {
	return &Currency{
		db: db,
	}
}

func (r *Currency) Create(data domains.Currency) (*domains.Currency, error) {
	currency := data.ToModel()
	if err := r.db.Create(&currency).Error; err != nil {
		return nil, err
	} else {
		return domains.FromCurrencyModel(currency), nil
	}
}

func (r *Currency) GetByBaseDate(base string, date time.Time) (*domains.Currency, error) {
	var currency models.Currency
	if err := r.db.Where("base = ? AND date = ?", base, date).First(&currency).Error; err != nil {
		return nil, err
	} else {
		return domains.FromCurrencyModel(&currency), nil
	}
}

func (r *Currency) GetLatest() (*domains.Currency, error) {
	var currency models.Currency
	if err := r.db.Order("date desc").First(&currency).Error; err != nil {
		return nil, err
	} else {
		return domains.FromCurrencyModel(&currency), nil
	}
}
