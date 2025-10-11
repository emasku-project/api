package cron

import (
	"fmt"

	"api/internal/features/general/domains"
	"api/internal/features/general/repositories"
	"api/pkg/market"
)

type Gold struct {
	market   *market.Market
	goldRepo *repositories.Asset
}

func NewGold(
	market *market.Market,
	goldRepo *repositories.Asset,
) *Gold {
	return &Gold{
		market:   market,
		goldRepo: goldRepo,
	}
}

func (c *Gold) Start() {
	gold, err := c.market.GetGoldPrice()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if _, err := c.goldRepo.Create(
		domains.Asset{
			Name:      gold.Name,
			Price:     gold.Price,
			Symbol:    gold.Symbol,
			UpdatedAt: gold.UpdatedAt,
		},
	); err != nil {
		fmt.Println(err.Error())
	}
}
