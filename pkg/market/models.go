package market

import "time"

type XAU struct {
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Symbol    string    `json:"symbol"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Dollar struct {
	Amount float64            `json:"amount"`
	Base   string             `json:"base"`
	Date   string             `json:"date"`
	Rates  map[string]float64 `json:"rates"`
}
