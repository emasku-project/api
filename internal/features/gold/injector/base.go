package injector

import (
	generalRepo "api/internal/features/general/repositories"
	"api/internal/features/gold/handlers"
	"api/internal/features/gold/repositories"
	"api/internal/features/gold/services"
	"api/pkg/database"
	"github.com/google/wire"
)

type GoldHandlers struct {
	Gold *handlers.Gold
}

func NewGoldHandlers(
	gold *handlers.Gold,
) *GoldHandlers {
	return &GoldHandlers{
		Gold: gold,
	}
}

var (
	Set = wire.NewSet(
		handlers.NewGold,
		services.NewGold,
		repositories.NewGold,
		generalRepo.NewAsset,
		generalRepo.NewCurrency,
		database.New,
	)
)
