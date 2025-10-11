package handlers

import (
	"net/http"
	"strconv"

	"api/internal/features/gold/dto/requests"
	"api/internal/features/gold/services"
	"api/pkg/http/failure"
	"github.com/gin-gonic/gin"
)

type Gold struct {
	service *services.Gold
}

func NewGold(
	service *services.Gold,
) *Gold {
	return &Gold{
		service: service,
	}
}

// @id			CreateGold
// @tags		gold
// @param		body body requests.CreateGold true "body"
// @success		200 {object} responses.CreateGold
// @router		/api/v1/golds [post]
func (h *Gold) CreateGold(c *gin.Context) {
	var req requests.CreateGold
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if res, err := h.service.CreateGold(c, req); err != nil {
		c.AbortWithStatusJSON(
			err.Code, failure.Failure{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @id			GetAllGolds
// @tags		gold
// @success		200 {object} responses.GetAllGolds
// @router		/api/v1/golds [get]
func (h *Gold) GetAllGolds(c *gin.Context) {
	if res, err := h.service.GetAllGolds(c); err != nil {
		c.AbortWithStatusJSON(
			err.Code, failure.Failure{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @id			GetGoldById
// @tags		gold
// @param 		gold_id path int true "gold_id"
// @success		200 {object} responses.GetGoldById
// @router		/api/v1/golds/{gold_id} [get]
func (h *Gold) GetGoldById(c *gin.Context) {
	goldId, err := strconv.Atoi(c.Param("gold_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if res, err := h.service.GetGoldById(uint(goldId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, failure.Failure{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @id			DeleteGoldById
// @tags		gold
// @param 		gold_id path int true "gold_id"
// @success		200 {object} responses.DeleteGoldById
// @router		/api/v1/golds/{gold_id} [delete]
func (h *Gold) DeleteGoldById(c *gin.Context) {
	goldId, err := strconv.Atoi(c.Param("gold_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if res, err := h.service.DeleteGoldById(uint(goldId)); err != nil {
		c.AbortWithStatusJSON(
			err.Code, failure.Failure{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}
