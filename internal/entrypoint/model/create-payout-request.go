package model

type CreatePayoutRequest struct {
	ExternalID string  `json:"external_id"`
	Amount     float64 `json:"amount"`
	MerchantID string  `json:"merchant_id"`
	Country    string  `json:"country"`
	Currency   string  `json:"currency"`
}

func (c CreatePayoutRequest) GetExternalID() string {
	return c.ExternalID
}

func (c CreatePayoutRequest) GetAmount() float64 {
	return c.Amount
}

func (c CreatePayoutRequest) GetCountry() string {
	return c.Country
}

func (c CreatePayoutRequest) GetMerchantID() string {
	return c.MerchantID
}

func (c CreatePayoutRequest) GetCurrency() string {
	return c.Currency
}
