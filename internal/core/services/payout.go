package services

import (
	"context"
	"github.com/shopspring/decimal"
	"golang-example/internal/core/command"
	"golang-example/internal/core/domain"
	"golang-example/internal/core/ports"
	"golang-example/internal/errors"
	"golang-example/pkg/uidgen"
)

type PayoutService struct {
	repository ports.PayoutRepository
}

func NewPayoutService(repository ports.PayoutRepository) *PayoutService {
	return &PayoutService{repository: repository}
}

func (p PayoutService) Create(ctx context.Context, pr command.CreatePayout) error {
	//validate fields
	err := validateFields(pr)
	if err != nil {
		return err
	}

	//Save payout
	payout := domain.NewPayout(
		uidgen.New().NewID(),
		pr.GetExternalID(),
		pr.GetAmount(),
		pr.GetMerchantID(),
		pr.GetCountry(),
		pr.GetCurrency())

	return p.repository.Save(ctx, payout)
}

func (p PayoutService) Get(ctx context.Context, id string) (domain.Payout, error) {
	if id == "" {
		return domain.Payout{}, errors.NewInvalidParamError("invalid id", nil)
	}

	return p.Get(ctx, id)
}

func validateFields(request command.CreatePayout) error {
	if request.GetExternalID() == "" ||
		request.GetCountry() == "" ||
		request.GetMerchantID() == "" ||
		request.GetAmount().LessThanOrEqual(decimal.Zero) {
		return errors.NewInvalidParamError("algun campo esta mal", nil)
	}
	return nil
}
