package user

import (
	"api/internal/features/user/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func User(g *gin.RouterGroup) {
	handlers := injector.InitUserHandlers()

	routes.Auth(g, handlers.Auth)
}
