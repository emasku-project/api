package http

import (
	"api/internal/features/general"
	"api/internal/features/gold"
	"api/internal/features/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer() {
	router := gin.Default()

	// cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))

	v1 := router.Group("/api/v1")

	user.User(v1)
	gold.Gold(v1)
	general.General(v1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
