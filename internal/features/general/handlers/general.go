package handlers

import (
	"net/http"

	"api/internal/features/general/services"
	"api/pkg/http/failure"
	"github.com/gin-gonic/gin"
)

type General struct {
	service *services.General
}

func NewGeneral(
	service *services.General,
) *General {
	return &General{
		service: service,
	}
}

// @id			GetSummary
// @tags		general
// @success		200 {object} responses.GetSummary
// @router		/api/v1/general/summary [get]
func (h *General) GetSummary(c *gin.Context) {
	if res, err := h.service.GetSummary(c); err != nil {
		c.AbortWithStatusJSON(
			err.Code, failure.Failure{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @id			GetMarketSummary
// @tags		general
// @success		200 {object} responses.GetMarketSummary
// @router		/api/v1/general/market-summary [get]
func (h *General) GetMarketSummary(c *gin.Context) {
	if res, err := h.service.GetMarketSummary(c); err != nil {
		c.AbortWithStatusJSON(
			err.Code, failure.Failure{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @id			GetSettings
// @tags		general
// @success		200 {object} responses.GetSettings
// @router		/api/v1/general/settings [get]
func (h *General) GetSettings(c *gin.Context) {
	if res, err := h.service.GetSettings(c); err != nil {
		c.AbortWithStatusJSON(
			err.Code, failure.Failure{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}
