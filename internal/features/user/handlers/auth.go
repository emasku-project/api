package handlers

import (
	"net/http"

	"api/internal/features/user/dto/requests"
	"api/internal/features/user/services"
	"api/pkg/http/failure"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	service *services.Auth
}

func NewAuth(
	service *services.Auth,
) *Auth {
	return &Auth{
		service: service,
	}
}

// @id			Login
// @tags		auth
// @param 		body body requests.Login true "body"
// @success		200 {object} responses.Login
// @router		/api/v1/auth/login [post]
func (h *Auth) Login(c *gin.Context) {
	var req requests.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if res, err := h.service.Login(req); err != nil {
		c.AbortWithStatusJSON(
			err.Code, failure.Failure{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @id			Register
// @tags		auth
// @param 		body body requests.Register true "body"
// @success		200 {object} responses.Register
// @router		/api/v1/auth/register [post]
func (h *Auth) Register(c *gin.Context) {
	var req requests.Register
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if res, err := h.service.Register(req); err != nil {
		c.AbortWithStatusJSON(
			err.Code, failure.Failure{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @id			Logout
// @tags		auth
// @success		200 {object} responses.Logout
// @router		/api/v1/auth/logout [post]
func (h *Auth) Logout(c *gin.Context) {
	if res, err := h.service.Logout(c); err != nil {
		c.AbortWithStatusJSON(
			err.Code, failure.Failure{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}
