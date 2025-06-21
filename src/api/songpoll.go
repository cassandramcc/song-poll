package api

import (
	"context"
)

func SongPoll(ctx context.Context, dataPath string) {
	// client := StartServer()
	// artists := data.GetDataFromCSV(dataPath)
	// for _, artist := range artists {
	// 	fmt.Println("• Processing artist: ", artist.Name)
	// 	albums := GetArtistAlbumsAfterDate(ctx, client, artist.URI, artist.LastVisted)
	// 	for _, album := range albums {
	// 		fmt.Println("  • Adding tracks from album: ", album.Name)
	// 		tracks := GetAlbumTracks(ctx, client, album.ID)
	// 		client.AddTracksToPlaylist(ctx, "4Eld3RdjPxPsiZaIX1Q0ID", GetTrackIDs(tracks)...)
	// 		fmt.Println("  ✓ Tracks added to playlist")
	// 	}
	// 	fmt.Println("✓ ", artist.Name, " processed")
	// }
}
