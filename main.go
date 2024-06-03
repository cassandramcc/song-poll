package main

import (
	"context"

	"github.com/cassandramcc/songpoll/api"
)

func main() {
	poll()
}

func poll() {
	api.SongPoll(context.Background())
}
