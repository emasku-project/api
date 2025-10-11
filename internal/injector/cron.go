//go:build wireinject
// +build wireinject

package injector

import (
	"api/internal/cron"
	"api/internal/features/general/repositories"
	"api/pkg/database"
	"api/pkg/market"
	"github.com/google/wire"
)

func InitGoldCron() *cron.Gold {
	wire.Build(
		cron.NewGold,
		market.New,
		repositories.NewAsset,
		database.New,
	)
	return nil
}

func InitDollarCron() *cron.Dollar {
	wire.Build(
		cron.NewDollar,
		market.New,
		repositories.NewCurrency,
		database.New,
	)
	return nil
}
