package command

type CreatePayout interface {
	GetExternalID() string
	GetAmount() float64
	GetCountry() string
	GetMerchantID() string
	GetCurrency() string
}

type GetPayout interface {
	GetID() string
}
