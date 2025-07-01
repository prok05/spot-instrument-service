package market

import (
	"context"
	"fmt"
	"github.com/prok05/spot-instrument-service/internal/entity"
	"github.com/prok05/spot-instrument-service/internal/repo"
)

type UseCase struct {
	repo repo.MarketRepo
}

func New(r repo.MarketRepo) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) ViewMarkets(ctx context.Context, in entity.ViewMarketsRequest) (entity.ViewMarketsResponse, error) {
	markets, err := uc.repo.GetActive(ctx)
	if err != nil {
		return entity.ViewMarketsResponse{}, fmt.Errorf("MarketUseCase - ViewMarkets - uc.repo.GetActive: %w", err)
	}

	return entity.ViewMarketsResponse{Markets: markets}, nil
}
