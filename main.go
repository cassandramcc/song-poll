package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cassandramcc/songpoll/api"
	"github.com/cassandramcc/songpoll/config"
	"github.com/cassandramcc/songpoll/model"
	"github.com/cassandramcc/songpoll/mongo"
)

func main() {

	// args := os.Args[1:]

	// if len(args) < 1 {
	// 	fmt.Println("provide a path to the data")
	// 	os.Exit(1)
	// }

	// dataPath := os.Args[1]
	// poll(dataPath)
	cfg := config.CreateConfig()
	ctx := context.Background()
	client, err := mongo.CreateClient(ctx, cfg.MongoConfig)
	if err != nil {
		panic(err)
	}
	err = client.CreateArtistURIUniqueIndex(ctx)
	if err != nil {
		panic(err)
	}

	err = client.AddArtist(ctx, model.Artist{
		URI:        "test",
		Name:       "Juche",
		LastVisted: time.Now(),
	})
	if err != nil {
		panic(err)
	}

	artsits, err := client.GetArtists(ctx, 10)
	if err != nil {
		panic(err)
	}

	for _, a := range artsits {
		fmt.Println(a.Name)
	}
}

func poll(dataPath string) {
	api.SongPoll(context.Background(), dataPath)
}
