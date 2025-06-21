package main

import (
	"context"
	"log"

	"github.com/cassandramcc/songpoll/config"
	"github.com/cassandramcc/songpoll/core"
	"github.com/cassandramcc/songpoll/mongo"
	"github.com/cassandramcc/songpoll/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.CreateConfig()
	client, err := mongo.CreateClient(context.Background(), cfg.MongoConfig)
	if err != nil {
		panic(err)
	}
	poller := core.NewPoller(client)
	server.StartServer(poller)
}
