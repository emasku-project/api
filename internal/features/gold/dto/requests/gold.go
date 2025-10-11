package requests

type CreateGold struct {
	PurchaseDate string  `json:"purchase_date" validate:"required"`
	Weight       float64 `json:"weight" validate:"required"`
	Price        float64 `json:"price" validate:"required"`
	Note         string  `json:"note" validate:"required"`
} // @name CreateGoldReq
