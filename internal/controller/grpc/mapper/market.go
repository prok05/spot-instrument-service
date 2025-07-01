package mapper

import (
	v1 "github.com/prok05/spot-instrument-service/api/proto/v1/gen"
	"github.com/prok05/spot-instrument-service/internal/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromProtoViewMarketsRequest(pb *v1.ViewMarketsRequest) (entity.ViewMarketsRequest, error) {
	roles := make([]string, 0, len(pb.UserRoles))

	for _, r := range pb.UserRoles {
		roles = append(roles, string(r))
	}

	return entity.ViewMarketsRequest{
		UserRoles: roles,
	}, nil
}

func ToProtoViewMarketsResponse(in entity.ViewMarketsResponse) *v1.ViewMarketsResponse {
	return &v1.ViewMarketsResponse{
		Markets: toProtoMarkets(in.Markets),
	}
}

func toProtoMarkets(ms []entity.Market) []*v1.Market {
	protoMarkets := make([]*v1.Market, 0, len(ms))

	for _, m := range ms {
		var t *timestamppb.Timestamp

		if !m.DeletedAt.IsZero() {
			t = timestamppb.New(m.DeletedAt)
		}

		protoMarkets = append(protoMarkets, &v1.Market{
			Id:        m.ID.String(),
			Name:      m.Name,
			Enabled:   m.Enabled,
			DeletedAt: t,
		})
	}

	return protoMarkets
}
