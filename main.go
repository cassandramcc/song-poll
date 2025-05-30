package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cassandramcc/songpoll/api"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("provide a path to the data")
		os.Exit(1)
	}

	dataPath := os.Args[1]
	poll(dataPath)
}

func poll(dataPath string) {
	api.SongPoll(context.Background(), dataPath)
}
