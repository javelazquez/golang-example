package ports

import (
	"context"
	"golang-example/internal/core/command"
	"golang-example/internal/core/domain"
)

type PayoutService interface {
	Create(ctx context.Context, payout command.CreatePayout) error
	Get(ctx context.Context, id string) (domain.Payout, error)
}
