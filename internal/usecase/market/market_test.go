package market

import (
	"context"
	"github.com/google/uuid"
	"github.com/prok05/spot-instrument-service/internal/entity"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

type mockMarketRepo struct {
	mock.Mock
}

func (m *mockMarketRepo) GetActive(ctx context.Context) ([]entity.Market, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.Market), args.Error(1)
}

func TestUseCase_ViewMarkets(t *testing.T) {
	ctx := context.Background()
	repoMock := &mockMarketRepo{}
	uc := New(repoMock)

	input := entity.ViewMarketsRequest{}
	exprectedMarkets := []entity.Market{
		{ID: uuid.New(), Enabled: true},
		{ID: uuid.New(), Enabled: true},
	}

	repoMock.On("GetActive", mock.Anything).Return(exprectedMarkets, nil)

	result, err := uc.ViewMarkets(ctx, input)
	require.NoError(t, err)
	require.Equal(t, exprectedMarkets, result.Markets)

	repoMock.AssertExpectations(t)
}
