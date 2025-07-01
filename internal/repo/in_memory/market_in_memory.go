package in_memory

import (
	"context"
	"github.com/google/uuid"
	"github.com/prok05/spot-instrument-service/internal/entity"
	"sync"
	"time"
)

type InMemoryMarketRepo struct {
	store map[uuid.UUID]entity.Market
	mu    sync.RWMutex
}

func New() *InMemoryMarketRepo {
	initMarkets := []entity.Market{
		{ID: uuid.New(), Name: "Market 1", Enabled: true},
		{ID: uuid.New(), Name: "Market 2", Enabled: true},
		{ID: uuid.New(), Name: "Market 2", Enabled: false, DeletedAt: time.Now()},
	}

	repo := &InMemoryMarketRepo{
		store: make(map[uuid.UUID]entity.Market),
	}

	for _, m := range initMarkets {
		repo.store[m.ID] = m
	}

	return repo
}

func (s *InMemoryMarketRepo) GetActive(ctx context.Context) ([]entity.Market, error) {
	markets := make([]entity.Market, 0)

	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.store) == 0 {
		return markets, nil
	}

	for _, v := range s.store {
		if v.Enabled && v.DeletedAt.IsZero() {
			markets = append(markets, v)
		}
	}

	return markets, nil
}
