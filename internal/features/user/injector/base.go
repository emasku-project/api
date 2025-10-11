package injector

import (
	"api/internal/features/user/handlers"
	"api/internal/features/user/repositories"
	"api/internal/features/user/services"
	"api/pkg/database"
	"github.com/google/wire"
)

type UserHandlers struct {
	Auth *handlers.Auth
}

func NewUserHandlers(
	auth *handlers.Auth,
) *UserHandlers {
	return &UserHandlers{
		Auth: auth,
	}
}

var (
	Set = wire.NewSet(
		handlers.NewAuth,
		services.NewAuth,
		repositories.NewUser,
		repositories.NewSession,
		database.New,
	)
)
