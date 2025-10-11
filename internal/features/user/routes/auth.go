package routes

import (
	"api/internal/features/user/handlers"
	"api/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func Auth(g *gin.RouterGroup, handler *handlers.Auth) {
	relativePath := "/auth"
	{
		group := g.Group(relativePath)
		group.POST("/login", handler.Login)
		group.POST("/register", handler.Register)
	}
	{
		group := g.Group(relativePath).Use(middlewares.AuthMiddleware())
		group.POST("/logout", handler.Logout)
	}
}
