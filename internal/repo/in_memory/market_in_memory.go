package in_memory

import (
	"context"
	"github.com/google/uuid"
	"github.com/prok05/spot-instrument-service/internal/entity"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"sync"
	"time"
)

type InMemoryMarketRepo struct {
	store  map[uuid.UUID]entity.Market
	mu     sync.RWMutex
	tracer trace.Tracer
}

func New() *InMemoryMarketRepo {
	staticID, _ := uuid.Parse("211c3ead-82fd-4ae2-becd-c637baf5632f")

	initMarkets := []entity.Market{
		{ID: staticID, Name: "Market 1", Enabled: true},
		{ID: uuid.New(), Name: "Market 2", Enabled: true},
		{ID: uuid.New(), Name: "Market 2", Enabled: false, DeletedAt: time.Now()},
	}

	repo := &InMemoryMarketRepo{
		store:  make(map[uuid.UUID]entity.Market),
		tracer: otel.Tracer("spot-instrument-service/in-memory-market-repo"),
	}

	for _, m := range initMarkets {
		repo.store[m.ID] = m
	}

	return repo
}

func (s *InMemoryMarketRepo) GetActive(ctx context.Context) ([]entity.Market, error) {
	ctx, span := s.tracer.Start(ctx, "GetActive")
	defer span.End()

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
