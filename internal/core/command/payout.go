package command

import "github.com/shopspring/decimal"

type CreatePayout interface {
	GetExternalID() string
	GetAmount() decimal.Decimal
	GetCountry() string
	GetMerchantID() string
	GetCurrency() string
}

type GetPayout interface {
	GetID() string
}
