package general

import (
	"api/internal/features/general/routes"
	"api/internal/injector"
	"github.com/gin-gonic/gin"
)

func General(g *gin.RouterGroup) {
	handlers := injector.InitGeneralHandlers()

	routes.General(g, handlers.General)
}
