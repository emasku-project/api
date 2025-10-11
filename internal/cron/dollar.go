package cron

import (
	"errors"
	"fmt"
	"time"

	"api/internal/features/general/domains"
	"api/internal/features/general/repositories"
	"api/pkg/market"
	"gorm.io/gorm"
)

type Dollar struct {
	market       *market.Market
	currencyRepo *repositories.Currency
}

func NewDollar(
	market *market.Market,
	currencyRepo *repositories.Currency,
) *Dollar {
	return &Dollar{
		market:       market,
		currencyRepo: currencyRepo,
	}
}

func (c *Dollar) Start() {
	dollar, err := c.market.GetDollarPrice()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dollarDate, err := time.Parse("2006-01-02", dollar.Date)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if _, err := c.currencyRepo.GetByBaseDate(dollar.Base, dollarDate); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if _, err := c.currencyRepo.Create(
				domains.Currency{
					Base: dollar.Base,
					Date: dollarDate,
					Rate: dollar.Rates["IDR"],
				},
			); err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
}
