package mongo

import (
	"context"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoClient struct {
	Client *mongo.Client
}

func CreateClient(ctx context.Context) (*MongoClient, error) {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to mongo")
	}
	return &MongoClient{
		Client: client,
	}, nil
}

func (m *MongoClient) Shutdown(ctx context.Context) error {
	err := m.Client.Disconnect(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to shutdown mongo client")
	}
	return nil
}

func (m *MongoClient) Artists() *mongo.Collection {
	return m.Client.Database("Song-Poll").Collection("artists")
}
