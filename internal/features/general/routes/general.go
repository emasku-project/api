package routes

import (
	"api/internal/features/general/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func General(g *gin.RouterGroup, handler *handlers.General) {
	relativePath := "/general"
	{
		group := g.Group(relativePath)
		group.GET("/market-summary", handler.GetMarketSummary)
	}
	{
		group := g.Group(relativePath).Use(middlewares.AuthMiddleware())
		group.GET("/summary", handler.GetSummary)
	}
}
