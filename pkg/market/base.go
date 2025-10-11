package market

import (
	"encoding/json"
	"net/http"
)

type Market struct{}

func New() *Market {
	return &Market{}
}

func (m *Market) GetGoldPrice() (*XAU, error) {
	xauRes, err := http.Get("https://api.gold-api.com/price/XAU")
	if err != nil {
		return nil, err
	}
	defer xauRes.Body.Close()

	var xau XAU
	if err := json.NewDecoder(xauRes.Body).Decode(&xau); err != nil {
		return nil, err
	} else {
		return &xau, nil
	}
}

func (m *Market) GetDollarPrice() (*Dollar, error) {
	idrRes, err := http.Get("https://api.frankfurter.dev/v1/latest?base=USD&symbols=IDR")
	if err != nil {
		return nil, err
	}
	defer idrRes.Body.Close()

	var dollar Dollar
	if err := json.NewDecoder(idrRes.Body).Decode(&dollar); err != nil {
		return nil, err
	} else {
		return &dollar, nil
	}
}

func (m *Market) GetGoldPricePerGram(percent float64) (float64, error) {
	if percent == 0 {
		percent = 100
	}

	xau, err := m.GetGoldPrice()
	if err != nil {
		return 0, err
	}

	dollar, err := m.GetDollarPrice()
	if err != nil {
		return 0, err
	}

	return (xau.Price * dollar.Rates["IDR"]) / 28.34952 * (percent / 100), nil
}
