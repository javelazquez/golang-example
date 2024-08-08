package domain

import "github.com/shopspring/decimal"

type Payout struct {
	Id         string
	ExternalID string
	Amount     decimal.Decimal
	MerchantID string
	Country    string
	Currency   string
	Status     string
}

func NewPayout(id string, externalID string, amount decimal.Decimal, merchantID string, country string, currency string) Payout {
	return Payout{
		Id:         id,
		ExternalID: externalID,
		Amount:     amount,
		MerchantID: merchantID,
		Country:    country,
		Currency:   currency,
		Status:     "created"}
}
