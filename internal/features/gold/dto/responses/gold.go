package responses

import "api/internal/features/gold/domains"

type CreateGold struct {
	Gold domains.Gold `json:"gold" validate:"required"`
} // @name CreateGoldRes

type GetAllGoldsItem struct {
	Gold           domains.Gold `json:"gold" validate:"required"`
	DurationInDays int          `json:"duration_in_days" validate:"required"`
	SellPrice      float64      `json:"sell_price" validate:"required"`
	Profit         float64      `json:"profit" validate:"required"`
}

type GetAllGolds struct {
	Items []GetAllGoldsItem `json:"items" validate:"required"`
} // @name GetAllGoldsRes

type GetGoldById struct {
	Gold domains.Gold `json:"gold" validate:"required"`
} // @name GetGoldByIdRes

type DeleteGoldById struct {
	Message string `json:"message" validate:"required"`
} // @name DeleteGoldByIdRes
