package services

import (
	"api/internal/features/general/dto/responses"
	"api/internal/features/general/repositories"
	goldRepo "api/internal/features/gold/repositories"
	"api/pkg/http/failure"
	"api/pkg/utils"
	"github.com/gin-gonic/gin"
)

type General struct {
	assetRepo    *repositories.Asset
	currencyRepo *repositories.Currency
	goldRepo     *goldRepo.Gold
}

func NewGeneral(
	assetRepo *repositories.Asset,
	currencyRepo *repositories.Currency,
	goldRepo *goldRepo.Gold,
) *General {
	return &General{
		assetRepo:    assetRepo,
		currencyRepo: currencyRepo,
		goldRepo:     goldRepo,
	}
}

func (s *General) GetSummary(c *gin.Context) (*responses.GetSummary, *failure.App) {
	session, err := utils.GetAuthenticatedSession(c)
	if err != nil {
		return nil, failure.NewUnauthorized()
	}

	asset, err := s.assetRepo.GetLatest()
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	currency, err := s.currencyRepo.GetLatest()
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	golds, err := s.goldRepo.GetAll(session.UserId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	jewPrice := asset.Price * currency.Rate / 31.1034767696 * .75

	totalBuyPrice := 0.0
	totalWeight := 0.0
	totalSellPrice := 0.0
	for _, v := range *golds {
		totalBuyPrice += v.Price
		totalWeight += v.Weight
		totalSellPrice += v.Weight * jewPrice
	}

	return &responses.GetSummary{
		TotalBuyPrice:  totalBuyPrice,
		TotalWeight:    totalWeight,
		TotalSellPrice: totalSellPrice,
		TotalProfit:    totalSellPrice - totalBuyPrice,
	}, nil
}

func (s *General) GetMarketSummary() (*responses.GetMarketSummary, *failure.App) {
	asset, err := s.assetRepo.GetLatest()
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	currency, err := s.currencyRepo.GetLatest()
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	return &responses.GetMarketSummary{
		GlobalXAUPrice:      asset.Price,
		GlobalXAUUpdatedAt:  asset.UpdatedAt,
		DollarRate:          currency.Rate,
		DollarUpdatedAt:     currency.Date,
		XAUPriceOunce:       asset.Price * currency.Rate,
		XAUPriceGram:        asset.Price * currency.Rate / 31.1034767696,
		XAUJewelryPriceGram: asset.Price * currency.Rate / 31.1034767696 * .75,
	}, nil
}
