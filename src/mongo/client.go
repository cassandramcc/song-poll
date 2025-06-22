package mongo

import (
	"context"
	"fmt"

	"github.com/cassandramcc/songpoll/src/config"
	"github.com/cassandramcc/songpoll/src/model"
	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoClient struct {
	artists *mongo.Collection
	client  *mongo.Client
}

func CreateClient(ctx context.Context, cfg *config.MongoConfig) (*MongoClient, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to mongo")
	}
	return &MongoClient{
		artists: client.Database(cfg.Database).Collection(cfg.ArtistsCollection),
		client:  client,
	}, nil
}

func (m *MongoClient) Shutdown(ctx context.Context) error {
	err := m.client.Disconnect(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to shutdown mongo client")
	}
	return nil
}

func (m *MongoClient) CreateArtistURIUniqueIndex(ctx context.Context) error {
	_, err := m.artists.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.M{"uri": 1},
			Options: options.Index().SetUnique(true),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create index in mongo")
	}

	return nil
}

func (m *MongoClient) AddArtist(ctx context.Context, artist *model.Artist) error {
	_, err := m.artists.InsertOne(ctx, artist)
	if err != nil {
		return errors.Wrap(err, "failed to insert artist into mongo")
	}
	fmt.Println("Artist", artist.Name, "added")
	return nil
}

func (m *MongoClient) GetArtists(ctx context.Context, pageSize int) ([]*model.Artist, error) {
	cursor, err := m.artists.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get artists")
	}

	artists := []*model.Artist{}
	for cursor.Next(context.TODO()) {
		var result *model.Artist
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println("❌ failed to decode", cursor.Current)
		}
		artists = append(artists, result)
	}

	return artists, nil
}
