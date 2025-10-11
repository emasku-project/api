package responses

import "time"

type Price struct {
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Symbol    string    `json:"symbol"`
	UpdatedAt time.Time `json:"updatedAt"`
}
