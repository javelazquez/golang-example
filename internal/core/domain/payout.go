package domain

// Definición de un nuevo tipo para los estados de Payout
type PayoutStatus string

// Constantes para los posibles estados
const (
	StatusCreated   PayoutStatus = "CREATED"
	StatusPending   PayoutStatus = "PENDING"
	StatusCompleted PayoutStatus = "COMPLETED"
	StatusRejected  PayoutStatus = "REJECTED"
)

type Payout struct {
	Id         string       `dynamodbav:"id" json:"id"`
	ExternalID string       `dynamodbav:"external_id" json:"external_id"`
	Amount     float64      `dynamodbav:"amount" json:"amount"`
	MerchantID string       `dynamodbav:"merchant_id" json:"merchant_id"`
	Country    string       `dynamodbav:"country" json:"country"`
	Currency   string       `dynamodbav:"currency" json:"currency"`
	Status     PayoutStatus `dynamodbav:"status" json:"status"`
}

func NewPayout(id string, externalID string, amount float64, merchantID string, country string, currency string) Payout {
	return Payout{
		Id:         id,
		ExternalID: externalID,
		Amount:     amount,
		MerchantID: merchantID,
		Country:    country,
		Currency:   currency,
		Status:     StatusCreated, // Usar la constante para establecer el estado inicial
	}
}

func (p Payout) GetStatus() string {
	return string(p.Status)
}
