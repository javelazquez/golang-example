package adapters

import (
	"context"
	"golang-example/internal/core/domain"
	"golang-example/pkg/kvs"
)

type payoutRepository struct {
	storage kvs.KVS
}

func NewPayoutRepository(storage kvs.KVS) *payoutRepository {
	return &payoutRepository{storage: storage}
}

func (p payoutRepository) Save(ctx context.Context, payout domain.Payout) error {
	return p.storage.Save(ctx, payout)
}

func (p payoutRepository) Get(ctx context.Context, id string) (domain.Payout, error) {
	payout := domain.Payout{}
	key := make(map[string]any)
	key["id"] = id
	err := p.storage.GetItem(ctx, key, &payout)
	if err != nil {
		return domain.Payout{}, err
	}
	return payout, nil
}
