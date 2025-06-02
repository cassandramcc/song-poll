package mongo

import (
	"context"
	"testing"

	"github.com/cassandramcc/songpoll/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type testVars struct {
	Client *MongoClient
}

func setup(t *testing.T) (*testVars, func()) {
	ctx := context.Background()
	cfg := config.CreateTestingConfig()
	client, err := CreateClient(ctx, cfg.MongoConfig)
	require.NoError(t, err)

	return &testVars{
			Client: client,
		}, func() {
			client.client.Database("Song-Poll-Test").Drop(context.Background())
			client.Shutdown(context.Background())
		}
}

func TestClient(t *testing.T) {
	ctx := context.Background()
	cfg := config.CreateTestingConfig()
	client, err := CreateClient(ctx, cfg.MongoConfig)
	require.NoError(t, err)

	err = client.client.Ping(ctx, nil)
	require.NoError(t, err)

	err = client.Shutdown(ctx)
	require.NoError(t, err)
}

func TestIndexCreation(t *testing.T) {
	tv, cleanup := setup(t)
	defer cleanup()

	err := tv.Client.CreateArtistURIUniqueIndex(context.Background())
	require.NoError(t, err)

	cursor, err := tv.Client.artists.Indexes().List(context.Background())
	require.NoError(t, err)

	indexes := []bson.M{}
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		require.NoError(t, err)
		indexes = append(indexes, result)
	}

	assert.Len(t, indexes, 2)
}
