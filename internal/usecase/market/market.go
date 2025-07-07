package market

import (
	"context"
	"fmt"
	"github.com/prok05/spot-instrument-service/internal/entity"
	"github.com/prok05/spot-instrument-service/internal/repo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type UseCase struct {
	repo   repo.MarketRepo
	tracer trace.Tracer
}

func New(r repo.MarketRepo) *UseCase {
	return &UseCase{
		repo:   r,
		tracer: otel.Tracer("spot-instrument-service/market-usecase"),
	}
}

func (uc *UseCase) ViewMarkets(ctx context.Context, in entity.ViewMarketsRequest) (entity.ViewMarketsResponse, error) {
	ctx, span := uc.tracer.Start(ctx, "ViewMarkets")
	defer span.End()

	markets, err := uc.repo.GetActive(ctx)
	if err != nil {
		return entity.ViewMarketsResponse{}, fmt.Errorf("MarketUseCase - ViewMarkets - uc.repo.GetActive: %w", err)
	}

	return entity.ViewMarketsResponse{Markets: markets}, nil
}
