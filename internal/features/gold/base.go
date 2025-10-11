package gold

import (
	"api/internal/features/gold/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func Gold(g *gin.RouterGroup) {
	handlers := injector.InitGoldHandlers()

	routes.Gold(g, handlers.Gold)
}
