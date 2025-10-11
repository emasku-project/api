package repositories

import (
	"api/internal/features/gold/domains"
	"api/internal/features/gold/models"
	"gorm.io/gorm"
)

type Gold struct {
	db *gorm.DB
}

func NewGold(db *gorm.DB) *Gold {
	return &Gold{db: db}
}

func (r *Gold) Create(data domains.Gold) (*domains.Gold, error) {
	gold := data.ToModel()
	if err := r.db.Create(&gold).Error; err != nil {
		return nil, err
	} else {
		return domains.FromGoldModel(gold), nil
	}
}

func (r *Gold) Get(id uint) (*domains.Gold, error) {
	var gold models.Gold
	if err := r.db.Where("id = ?", id).First(&gold).Error; err != nil {
		return nil, err
	} else {
		return domains.FromGoldModel(&gold), nil
	}
}

func (r *Gold) GetAll(userId uint) (*[]domains.Gold, error) {
	var golds []models.Gold
	if err := r.db.Where("user_id = ?", userId).Find(&golds).Error; err != nil {
		return nil, err
	} else {
		result := make([]domains.Gold, len(golds))
		for i, v := range golds {
			result[i] = *domains.FromGoldModel(&v)
		}

		return &result, nil
	}
}

func (r *Gold) Update(id uint, data domains.Gold) (*domains.Gold, error) {
	gold := data.ToModel()
	if err := r.db.Where("id = ?", id).Updates(&gold).Error; err != nil {
		return nil, err
	} else {
		return domains.FromGoldModel(gold), nil
	}
}

func (r *Gold) Delete(id uint) error {
	return r.db.Where("id = ?", id).Unscoped().Delete(&models.Gold{}).Error
}
