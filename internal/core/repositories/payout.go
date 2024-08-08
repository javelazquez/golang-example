package repositories

import (
	"context"
	"golang-example/internal/core/domain"
	"golang-example/pkg/kvs"
)

type PayoutRepository struct {
	storage kvs.KVS
}

func NewPayoutRepository(storage kvs.KVS) *PayoutRepository {
	return &PayoutRepository{storage: storage}
}

func (p PayoutRepository) Save(ctx context.Context, payout domain.Payout) error {
	return p.storage.Save(ctx, payout)
}

func (p PayoutRepository) Get(ctx context.Context, id string) (domain.Payout, error) {
	payout := domain.Payout{}
	key := make(map[string]any)
	key["id"] = id
	err := p.storage.GetItem(ctx, key, &payout)
	if err != nil {
		return domain.Payout{}, err
	}
	return payout, nil
}
