package utils

import (
	"errors"

	"api/internal/features/user/domains"
	"github.com/gin-gonic/gin"
)

func GetAuthenticatedUser(c *gin.Context) (*domains.User, error) {
	if user, exists := c.Value("user").(domains.User); !exists {
		return nil, errors.New("unauthenticated user")
	} else {
		return &user, nil
	}
}

func GetAuthenticatedSession(c *gin.Context) (*domains.Session, error) {
	if session, exists := c.Value("session").(domains.Session); !exists {
		return nil, errors.New("unauthenticated user")
	} else {
		return &session, nil
	}
}
