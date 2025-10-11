package domains

import (
	"time"

	"api/internal/features/gold/models"
)

type Gold struct {
	Id           uint      `json:"id" validate:"required"`
	UserId       uint      `json:"user_id" validate:"required"`
	PurchaseDate time.Time `json:"purchase_date" validate:"required"`
	Weight       float64   `json:"weight" validate:"required"`
	Price        float64   `json:"price" validate:"required"`
	Note         string    `json:"note" validate:"required"`
}

func FromGoldModel(m *models.Gold) *Gold {
	return &Gold{
		Id:           m.ID,
		UserId:       m.UserId,
		PurchaseDate: m.PurchaseDate,
		Weight:       m.Weight,
		Price:        m.Price,
		Note:         m.Note,
	}
}

func (g *Gold) ToModel() *models.Gold {
	return &models.Gold{
		UserId:       g.UserId,
		PurchaseDate: g.PurchaseDate,
		Weight:       g.Weight,
		Price:        g.Price,
		Note:         g.Note,
	}
}
