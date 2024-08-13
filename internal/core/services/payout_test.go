package services_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-example/internal/core/domain"
	"golang-example/internal/core/services"
	"testing"
)

type MockRequest struct {
	ExternalID string  `json:"external_id"`
	Amount     float64 `json:"amount"`
	MerchantID string  `json:"merchant_id"`
	Country    string  `json:"country"`
	Currency   string  `json:"currency"`
}

func (c MockRequest) GetExternalID() string {
	return c.ExternalID
}

func (c MockRequest) GetAmount() float64 {
	return c.Amount
}

func (c MockRequest) GetCountry() string {
	return c.Country
}

func (c MockRequest) GetMerchantID() string {
	return c.MerchantID
}

func (c MockRequest) GetCurrency() string {
	return c.Currency
}

// Mock repository
type MockPayoutRepository struct {
	mock.Mock
}

func (m *MockPayoutRepository) Save(ctx context.Context, payout domain.Payout) error {
	args := m.Called(ctx, payout)
	return args.Error(0)
}

func (m *MockPayoutRepository) Get(ctx context.Context, id string) (domain.Payout, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Payout), args.Error(1)
}

// MockUIDGen es un mock de la interfaz UIDGen
type MockUIDGen struct {
	mock.Mock
}

func (m *MockUIDGen) NewID() string {
	args := m.Called()
	return args.String(0)
}

func TestPayoutService_Create_Success(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockPayoutRepository)
	//mockUIDGen := new(MockUIDGen)
	service := services.NewPayoutService(mockRepo)

	//payoutIdExpected := "payout-id-00"
	request := MockRequest{
		ExternalID: "external_test_1",
		Amount:     100,
		MerchantID: "merchant-4",
		Country:    "AR",
		Currency:   "ARS",
	}

	/*payout := domain.Payout{
		Id:         payoutIdExpected,
		ExternalID: request.GetExternalID(),
		Amount:     request.GetAmount(),
		MerchantID: request.GetMerchantID(),
		Country:    request.GetCountry(),
		Currency:   request.GetCurrency(),
		Status:     "CREATED",
	}*/

	//mockUIDGen.On("NewID").Return(payoutIdExpected)
	mockRepo.On("Save", ctx, mock.Anything).Return(nil)

	_, err := service.Create(ctx, request)

	mockRepo.AssertCalled(t, "Save", ctx, mock.Anything)
	//mockUIDGen.AssertCalled(t, "New")
	mockRepo.AssertExpectations(t)
	//mockUIDGen.AssertExpectations(t)

	assert.NoError(t, err)
	//assert.Equal(t, payoutIdExpected, payoutID)
}
