package main

import (
	"context"

	"github.com/cassandramcc/songpoll/api"
)

func main() {
	poll()
}

func poll() {
	client := api.StartServer()
	tracks := api.GetAlbumTracks(context.Background(), client, "7rhCjq7EaXmQT1sTs8Ls01")
	client.AddTracksToPlaylist(context.Background(), "4Eld3RdjPxPsiZaIX1Q0ID", api.GetTrackIDs(tracks)...)
}
