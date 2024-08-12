package model

type GetPayoutRequest struct {
	Id string `json:"id"`
}

func NewGetPayoutRequest(id string) GetPayoutRequest {
	return GetPayoutRequest{Id: id}
}

func (g GetPayoutRequest) GetID() string {
	return g.Id
}
