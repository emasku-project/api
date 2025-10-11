package services

import (
	"time"

	generalRepo "api/internal/features/general/repositories"
	"api/internal/features/gold/domains"
	"api/internal/features/gold/dto/requests"
	"api/internal/features/gold/dto/responses"
	"api/internal/features/gold/repositories"
	"api/pkg/http/failure"
	"api/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Gold struct {
	goldRepo     *repositories.Gold
	assetRepo    *generalRepo.Asset
	currencyRepo *generalRepo.Currency
}

func NewGold(
	goldRepo *repositories.Gold,
	assetRepo *generalRepo.Asset,
	currencyRepo *generalRepo.Currency,
) *Gold {
	return &Gold{
		goldRepo:     goldRepo,
		assetRepo:    assetRepo,
		currencyRepo: currencyRepo,
	}
}

func (s *Gold) CreateGold(c *gin.Context, req requests.CreateGold) (
	*responses.CreateGold, *failure.App,
) {
	session, err := utils.GetAuthenticatedSession(c)
	if err != nil {
		return nil, failure.NewUnauthorized()
	}

	purchaseDate, err := utils.GetParsedDate(req.PurchaseDate)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	if res, err := s.goldRepo.Create(
		domains.Gold{
			UserId:       session.UserId,
			PurchaseDate: *purchaseDate,
			Weight:       req.Weight,
			Price:        req.Price,
			Note:         req.Note,
		},
	); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.CreateGold{
			Gold: *res,
		}, nil
	}
}

func (s *Gold) GetAllGolds(c *gin.Context) (*responses.GetAllGolds, *failure.App) {
	session, err := utils.GetAuthenticatedSession(c)
	if err != nil {
		return nil, failure.NewUnauthorized()
	}

	gold, err := s.assetRepo.GetLatest()
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	dollar, err := s.currencyRepo.GetLatest()
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	golds, err := s.goldRepo.GetAll(session.UserId)
	if err != nil {
		return nil, failure.NewInternal(err)
	}

	items := make([]responses.GetAllGoldsItem, len(*golds))
	for i, v := range *golds {
		sellPrice := (gold.Price * dollar.Rate / 28.34952 * .75) * v.Weight
		profit := sellPrice - v.Price

		items[i] = responses.GetAllGoldsItem{
			Gold:           v,
			DurationInDays: time.Now().Day() - v.PurchaseDate.Day(),
			SellPrice:      sellPrice,
			Profit:         profit,
		}
	}

	return &responses.GetAllGolds{
		Items: items,
	}, nil
}

func (s *Gold) GetGoldById(goldId uint) (*responses.GetGoldById, *failure.App) {
	if res, err := s.goldRepo.Get(goldId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.GetGoldById{
			Gold: *res,
		}, nil
	}
}

func (s *Gold) DeleteGoldById(goldId uint) (*responses.DeleteGoldById, *failure.App) {
	if err := s.goldRepo.Delete(goldId); err != nil {
		return nil, failure.NewInternal(err)
	} else {
		return &responses.DeleteGoldById{
			Message: "oke",
		}, nil
	}
}
