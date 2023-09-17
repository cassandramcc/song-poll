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
	//dateStrubf := "2023-01-01"
	//date, _ := time.Parse("2006-01-02", dateStrubf)
	//api.GetArtistAlbumsAfterDate(context.Background(), client, "3GZ8Rfap7VxAOzABiZEXCL", date)
	tracks := api.GetAlbumTracks(context.Background(), client, "7rhCjq7EaXmQT1sTs8Ls01")
	client.AddTracksToPlaylist(context.Background(), "4Eld3RdjPxPsiZaIX1Q0ID", api.GetTrackIDs(tracks)...)
}
