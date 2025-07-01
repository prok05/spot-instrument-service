package usecase

import (
	"context"
	"github.com/prok05/spot-instrument-service/internal/entity"
)

type (
	Market interface {
		ViewMarkets(ctx context.Context, in entity.ViewMarketsRequest) (entity.ViewMarketsResponse, error)
	}
)
