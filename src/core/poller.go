package core

import (
	"context"

	"github.com/cassandramcc/songpoll/src/model"
	"github.com/cassandramcc/songpoll/src/mongo"
	"github.com/pkg/errors"
)

type Poller struct {
	client *mongo.MongoClient
}

func (p *Poller) GetArtists() ([]*model.Artist, error) {
	artists, err := p.client.GetArtists(context.Background(), 10)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get artists from client")
	}

	return artists, nil
}

func NewPoller(client *mongo.MongoClient) *Poller {
	return &Poller{
		client: client,
	}
}
