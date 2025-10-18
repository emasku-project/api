package handlers

import (
	"net/http"

	"api/internal/features/general/dto/requests"
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

// @id			UpdateTaxSetting
// @tags		general
// @param 		body body requests.UpdateTaxSetting true "body"
// @success		200 {object} responses.UpdateTaxSetting
// @router		/api/v1/general/settings/tax [post]
func (h *General) UpdateTaxSetting(c *gin.Context) {
	var req requests.UpdateTaxSetting
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if res, err := h.service.UpdateTaxSetting(c, req); err != nil {
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
