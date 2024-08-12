package services

import (
	"context"
	"fmt"
	"golang-example/internal/appErrors"
	"golang-example/internal/core/command"
	"golang-example/internal/core/domain"
	"golang-example/internal/core/ports"
	"golang-example/pkg/uidgen"
)

type payoutService struct {
	repository ports.PayoutRepository
}

func NewPayoutService(repository ports.PayoutRepository) *payoutService {
	return &payoutService{repository: repository}
}

func (p payoutService) Create(ctx context.Context, pr command.CreatePayout) (string, error) {
	//validate fields
	err := validateFields(pr)
	if err != nil {
		return "", err
	}

	//SaveConditional payout
	payout := domain.NewPayout(
		uidgen.New().NewID(),
		pr.GetExternalID(),
		pr.GetAmount(),
		pr.GetMerchantID(),
		pr.GetCountry(),
		pr.GetCurrency())

	err = p.repository.Save(ctx, payout)
	if err != nil {
		return "", err
	}

	return payout.Id, nil
}

func (p payoutService) Get(ctx context.Context, request command.GetPayout) (domain.Payout, error) {
	if request.GetID() == "" {
		return domain.Payout{}, appErrors.NewIdError("id is required", nil)
	}

	payout, err := p.repository.Get(ctx, request.GetID())
	if err != nil {
		return domain.Payout{}, err
	}

	if payout.Id == "" {
		return domain.Payout{}, appErrors.NewNotFoundError(fmt.Sprintf("payout %s not found", request.GetID()))
	}

	return payout, nil
}

func validateFields(request command.CreatePayout) error {
	if request.GetExternalID() == "" {
		return appErrors.NewExternalIDError("ExternalID no puede estar vacío", nil)
	}
	if request.GetCountry() == "" {
		return appErrors.NewCountryError("Country no puede estar vacío", nil)
	}
	if request.GetMerchantID() == "" {
		return appErrors.NewMerchantIDError("MerchantID no puede estar vacío", nil)
	}
	if request.GetAmount() <= 0 {
		return appErrors.NewAmountError("Amount debe ser mayor que 0", nil)
	}
	if request.GetCurrency() == "" {
		return appErrors.NewCurrencyError("Currency no puede estar vacío", nil)
	}
	return nil
}
