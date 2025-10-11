package repositories

import (
	"api/internal/features/general/domains"
	"api/internal/features/general/models"
	"gorm.io/gorm"
)

type Asset struct {
	db *gorm.DB
}

func NewAsset(
	db *gorm.DB,
) *Asset {
	return &Asset{
		db: db,
	}
}

func (r *Asset) Create(data domains.Asset) (*domains.Asset, error) {
	asset := data.ToModel()
	if err := r.db.Create(&asset).Error; err != nil {
		return nil, err
	} else {
		return domains.FromAssetModel(asset), nil
	}
}

func (r *Asset) GetLatest() (*domains.Asset, error) {
	var asset models.Asset
	if err := r.db.Order("updated_at desc").First(&asset).Error; err != nil {
		return nil, err
	} else {
		return domains.FromAssetModel(&asset), nil
	}
}
