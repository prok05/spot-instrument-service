package repo

import (
	"context"
	"github.com/prok05/spot-instrument-service/internal/entity"
)

type (
	MarketRepo interface {
		GetActive(ctx context.Context) ([]entity.Market, error)
	}
)
