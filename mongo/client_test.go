package mongo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	ctx := context.Background()
	client, err := CreateClient(ctx)
	require.NoError(t, err)

	err = client.Client.Ping(ctx, nil)
	require.NoError(t, err)

	err = client.Shutdown(ctx)
	require.NoError(t, err)
}
