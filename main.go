package main

import (
	"context"
	"time"

	"github.com/cassandramcc/songpoll/api"
)

func main() {
	client := api.StartServer()
	dateStrubf := "2023-01-01"
	date, _ := time.Parse("2006-01-02", dateStrubf)
	api.GetArtistAlbumsAfterDate(context.Background(), client, "3GZ8Rfap7VxAOzABiZEXCL", date)
}
