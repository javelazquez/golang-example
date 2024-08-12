package model

type CreatePayoutResponse struct {
	Id string `json:"id"`
}

func NewCreatePayoutResponse(id string) CreatePayoutResponse {
	return CreatePayoutResponse{Id: id}
}
