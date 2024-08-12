package ports

import (
	"context"
	"golang-example/internal/core/domain"
)

type PayoutRepository interface {
	Save(ctx context.Context, payout domain.Payout) error
	Get(ctx context.Context, id string) (domain.Payout, error)
}
