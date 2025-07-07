package v1_test

import (
	"context"
	"github.com/google/uuid"
	v1 "github.com/prok05/spot-instrument-service/api/proto/v1/gen"
	"github.com/prok05/spot-instrument-service/internal/controller/grpc"
	"github.com/prok05/spot-instrument-service/internal/controller/grpc/mapper"
	"github.com/prok05/spot-instrument-service/internal/entity"
	pkgLogger "github.com/prok05/spot-instrument-service/pkg/logger"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	pbgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"testing"
)

type mockMarketUseCase struct {
	mock.Mock
}

func (m *mockMarketUseCase) ViewMarkets(ctx context.Context, in entity.ViewMarketsRequest) (entity.ViewMarketsResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(entity.ViewMarketsResponse), args.Error(1)
}

func TestGRPC_ViewMarkets(t *testing.T) {
	ctx := context.Background()
	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	uc := &mockMarketUseCase{}
	logger, _ := pkgLogger.New("debug")
	srv := pbgrpc.NewServer()
	grpc.NewRouter(srv, uc, logger)

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("srv.Serve %v", err)
		}
	}()

	t.Cleanup(func() {
		srv.Stop()
	})

	t.Run("should return active markets", func(t *testing.T) {
		want := entity.ViewMarketsResponse{
			Markets: []entity.Market{
				{ID: uuid.New(), Name: "Market1", Enabled: true},
				{ID: uuid.New(), Name: "Market2", Enabled: true},
				{ID: uuid.New(), Name: "Market3", Enabled: true},
			},
		}
		uc.On("ViewMarkets", mock.Anything, entity.ViewMarketsRequest{UserRoles: make([]string, 0)}).Return(
			want, nil)

		dialer := func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}

		conn, err := pbgrpc.NewClient("passthrough:///", pbgrpc.WithContextDialer(dialer), pbgrpc.WithTransportCredentials(insecure.NewCredentials()))
		t.Cleanup(func() {
			conn.Close()
		})
		if err != nil {
			t.Fatalf("grpc.DialContext %v", err)
		}
		client := v1.NewSpotInstrumentServiceClient(conn)

		resp, err := client.ViewMarkets(ctx, &v1.ViewMarketsRequest{})
		require.NoError(t, err)

		mappedWant := mapper.ToProtoViewMarketsResponse(want)

		require.True(t, proto.Equal(resp, mappedWant))

		uc.AssertExpectations(t)
	})

}
