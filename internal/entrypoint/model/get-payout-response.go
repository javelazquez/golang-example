package model

import "golang-example/internal/core/domain"

type GetPayoutResponse struct {
	Id         string  `json:"id"`
	ExternalID string  `json:"external_id"`
	Amount     float64 `json:"amount"`
	MerchantID string  `json:"merchant_id"`
	Country    string  `json:"country"`
	Currency   string  `json:"currency"`
	Status     string  `json:"status"`
}

func NewGetPayoutResponse(payout domain.Payout) *GetPayoutResponse {
	return &GetPayoutResponse{
		Id:         payout.Id,
		ExternalID: payout.ExternalID,
		Amount:     payout.Amount,
		MerchantID: payout.MerchantID,
		Country:    payout.Country,
		Currency:   payout.Currency,
		Status:     payout.GetStatus(),
	}
}
