package in_memory

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInMemoryOrderRepo_GetActive(t *testing.T) {
	ctx := context.Background()
	repo := New()

	markets, err := repo.GetActive(ctx)
	require.NoError(t, err)

	for _, m := range markets {
		require.True(t, m.Enabled)
		require.Zero(t, m.DeletedAt)
	}
}
