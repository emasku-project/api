package injector

import (
	"api/internal/features/general/handlers"
	"api/internal/features/general/repositories"
	"api/internal/features/general/services"
	goldRepo "api/internal/features/gold/repositories"
	"api/pkg/database"
	"github.com/google/wire"
)

type GeneralHandlers struct {
	General *handlers.General
}

func NewGeneralHandlers(
	general *handlers.General,
) *GeneralHandlers {
	return &GeneralHandlers{
		General: general,
	}
}

var (
	Set = wire.NewSet(
		handlers.NewGeneral,
		services.NewGeneral,
		repositories.NewAsset,
		repositories.NewCurrency,
		goldRepo.NewGold,
		database.New,
	)
)
