package responses

import "time"

type GetSummary struct {
	TotalBuyPrice  float64 `json:"total_buy_price" validate:"required"`
	TotalWeight    float64 `json:"total_weight" validate:"required"`
	TotalSellPrice float64 `json:"total_sell_price" validate:"required"`
	TotalProfit    float64 `json:"total_profit" validate:"required"`
} // @name GetSummaryRes

type GetMarketSummary struct {
	GlobalXAUPrice      float64   `json:"global_xau_price"`
	GlobalXAUUpdatedAt  time.Time `json:"global_xau_updated_at"`
	DollarRate          float64   `json:"dollar_rate"`
	DollarUpdatedAt     time.Time `json:"dollar_updated_at"`
	XAUPriceGram        float64   `json:"xau_price_gram"`
	XAUJewelryPriceGram float64   `json:"xau_jewelry_price_gram"`
	TaxPercentage       float64   `json:"tax_percentage"`
} // @name GetMarketSummaryRes

type UpdateTaxSetting struct {
	Message string `json:"message"`
} // @name UpdateTaxSettingRes

type GetSettings struct {
	TaxPercentage float64 `json:"tax_percentage"`
} // @name GetSettingsRes
