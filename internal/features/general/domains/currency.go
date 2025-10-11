package domains

import (
	"time"

	"api/internal/features/general/models"
)

type Currency struct {
	Id   uint      `json:"id"`
	Base string    `json:"base"`
	Date time.Time `json:"date"`
	Rate float64   `json:"rate"`
}

func (c *Currency) ToModel() *models.Currency {
	return &models.Currency{
		Base: c.Base,
		Date: c.Date,
		Rate: c.Rate,
	}
}

func FromCurrencyModel(m *models.Currency) *Currency {
	return &Currency{
		Id:   m.ID,
		Base: m.Base,
		Date: m.Date,
		Rate: m.Rate,
	}
}
