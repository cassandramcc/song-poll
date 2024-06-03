package api

import (
	"context"
	"fmt"

	"github.com/cassandramcc/songpoll/data"
)

func SongPoll(ctx context.Context) {
	client := StartServer()
	artists := data.GetDataFromCSV(`data\test.csv`)
	for _, artist := range artists {
		fmt.Println("• Processing artist: ", artist.Name)
		albums := GetArtistAlbumsAfterDate(ctx, client, artist.ID, artist.LastVisted)
		for _, album := range albums {
			fmt.Println("  • Adding tracks from album: ", album.Name)
			tracks := GetAlbumTracks(ctx, client, album.ID)
			client.AddTracksToPlaylist(ctx, "4Eld3RdjPxPsiZaIX1Q0ID", GetTrackIDs(tracks)...)
			fmt.Println("  ✓ Tracks added to playlist")
		}
		fmt.Println("✓ ", artist.Name, " processed")
	}
}
