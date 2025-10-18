package requests

type UpdateTaxSetting struct {
	TaxPercentage float64 `json:"tax_percentage" validate:"required"`
} // @name UpdateTaxSettingReq
