package v1

import (
	"github.com/go-playground/validator/v10"
	v1 "github.com/prok05/spot-instrument-service/api/proto/v1/gen"
	"github.com/prok05/spot-instrument-service/internal/usecase"
	"github.com/prok05/spot-instrument-service/pkg/logger"
	pbgrpc "google.golang.org/grpc"
)

func NewMarketRouter(app *pbgrpc.Server, uc usecase.Market, l logger.Interface) {
	r := &V1{
		l:  l,
		uc: uc,
		v:  validator.New(validator.WithRequiredStructEnabled()),
	}

	{
		v1.RegisterSpotInstrumentServiceServer(app, r)
	}
}
