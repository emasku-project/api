package routes

import (
	"api/internal/features/gold/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func Gold(g *gin.RouterGroup, handler *handlers.Gold) {
	group := g.Group("/golds").Use(middlewares.AuthMiddleware())

	group.POST("", handler.CreateGold)
	group.GET("", handler.GetAllGolds)
	group.GET("/:gold_id", handler.GetGoldById)
	group.DELETE("/:gold_id", handler.DeleteGoldById)
}
