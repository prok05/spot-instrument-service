package v1

import (
	"github.com/go-playground/validator/v10"
	v1 "github.com/prok05/spot-instrument-service/api/proto/v1/gen"
	"github.com/prok05/spot-instrument-service/internal/usecase"
	"github.com/prok05/spot-instrument-service/pkg/logger"
)

type V1 struct {
	v1.SpotInstrumentServiceServer
	uc usecase.Market
	l  logger.Interface
	v  *validator.Validate
}
