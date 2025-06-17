package main

import (
	"context"

	"github.com/cassandramcc/songpoll/config"
	"github.com/cassandramcc/songpoll/core"
	"github.com/cassandramcc/songpoll/mongo"
	"github.com/cassandramcc/songpoll/server"
)

func main() {
	cfg := config.CreateConfig()
	client, err := mongo.CreateClient(context.Background(), cfg.MongoConfig)
	if err != nil {
		panic(err)
	}
	poller := core.NewPoller(client)
	server.StartServer(poller)
}
