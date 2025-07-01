package v1

import (
	"context"
	v1 "github.com/prok05/spot-instrument-service/api/proto/v1/gen"
	"github.com/prok05/spot-instrument-service/internal/controller/grpc/mapper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *V1) ViewMarkets(ctx context.Context, in *v1.ViewMarketsRequest) (*v1.ViewMarketsResponse, error) {
	domainR, err := mapper.FromProtoViewMarketsRequest(in)
	if err != nil {
		r.l.Error("grpc - V1 - ViewMarkets - mapper.FromProtoViewMarketsRequest", "mapper error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := r.v.Struct(domainR); err != nil {
		r.l.Error("grpc - V1 - ViewMarkets - r.v.Struct", "validation error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	respDomain, err := r.uc.ViewMarkets(ctx, domainR)
	if err != nil {
		r.l.Error("grpc - V1 - GetOrderStatus - r.uc.ViewMarkets", "internal error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return mapper.ToProtoViewMarketsResponse(respDomain), nil
}
