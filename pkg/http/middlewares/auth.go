package middlewares

import (
	"net/http"

	"api/internal/features/user/domains"
	"api/internal/features/user/models"
	"api/pkg/database"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	db := database.New()

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := authHeader[7:]

		// mendapatkan session berdasarkan token
		var session models.Session
		if err := db.Where("token = ?", token).First(&session).Error; err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// mendapatkan user berdasarkan id dari session
		var user models.User
		if err := db.Where("id = ?", session.UserId).First(&user).Error; err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(
			"user", domains.User{
				Id:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			},
		)
		c.Set(
			"session", domains.Session{
				Id:     session.ID,
				UserId: session.UserId,
				Token:  session.Token,
			},
		)
		c.Next()
	}
}
