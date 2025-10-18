package routes

import (
	"api/internal/features/general/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func General(g *gin.RouterGroup, handler *handlers.General) {
	relativePath := "/general"
	{
		group := g.Group(relativePath).Use(middlewares.AuthMiddleware())
		group.GET("/market-summary", handler.GetMarketSummary)
		group.GET("/summary", handler.GetSummary)

		group.POST("/settings/tax", handler.UpdateTaxSetting)
		group.GET("/settings", handler.GetSettings)
	}
}
