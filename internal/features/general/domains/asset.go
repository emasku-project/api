package domains

import (
	"time"

	"api/internal/features/general/models"
)

type Asset struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Symbol    string    `json:"symbol"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Asset) ToModel() *models.Asset {
	return &models.Asset{
		Name:      a.Name,
		Price:     a.Price,
		Symbol:    a.Symbol,
		UpdatedAt: a.UpdatedAt,
	}
}

func FromAssetModel(m *models.Asset) *Asset {
	return &Asset{
		Id:        m.ID,
		Name:      m.Name,
		Price:     m.Price,
		Symbol:    m.Symbol,
		UpdatedAt: m.UpdatedAt,
	}
}
