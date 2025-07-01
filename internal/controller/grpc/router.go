package grpc

import (
	v1 "github.com/prok05/spot-instrument-service/internal/controller/grpc/v1"
	"github.com/prok05/spot-instrument-service/internal/usecase"
	"github.com/prok05/spot-instrument-service/pkg/logger"
	pbgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewRouter(app *pbgrpc.Server, uc usecase.Market, l logger.Interface) {
	{
		v1.NewMarketRouter(app, uc, l)
	}

	reflection.Register(app)
}
