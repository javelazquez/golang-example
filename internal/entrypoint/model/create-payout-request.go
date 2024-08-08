package model

import "github.com/shopspring/decimal"

type CreatePayoutRequest struct {
	ExternalID string          `json:"external_id"`
	Amount     decimal.Decimal `json:"amount"`
	MerchantID string          `json:"merchant_id"`
	Country    string          `json:"country"`
	Currency   string          `json:"currency"`
}

func (c CreatePayoutRequest) GetExternalID() string {
	return c.ExternalID
}

func (c CreatePayoutRequest) GetAmount() decimal.Decimal {
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
