package ports

import (
	"context"
	"golang-example/internal/core/command"
	"golang-example/internal/core/domain"
)

type PayoutService interface {
	Create(ctx context.Context, payout command.CreatePayout) (string, error)
	Get(ctx context.Context, request command.GetPayout) (domain.Payout, error)
}
