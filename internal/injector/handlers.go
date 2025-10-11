//go:build wireinject
// +build wireinject

package injector

import (
	general "api/internal/features/general/injector"
	gold "api/internal/features/gold/injector"
	user "api/internal/features/user/injector"
	"github.com/google/wire"
)

func InitGoldHandlers() *gold.GoldHandlers {
	wire.Build(
		gold.NewGoldHandlers,
		gold.Set,
	)
	return nil
}

func InitGeneralHandlers() *general.GeneralHandlers {
	wire.Build(
		general.NewGeneralHandlers,
		general.Set,
	)
	return nil
}

func InitUserHandlers() *user.UserHandlers {
	wire.Build(
		user.NewUserHandlers,
		user.Set,
	)
	return nil
}
