package api

import (
	"context"
	"fmt"
	"time"

	"github.com/cassandramcc/songpoll/model"
	"github.com/zmb3/spotify/v2"
)

func GetArtistAlbums(ctx context.Context, client *spotify.Client, artistID string) []*model.Album {
	var results []*model.Album
	offset := 0
	complete := false
	for !complete {
		albumTypes := []spotify.AlbumType{spotify.AlbumTypeAlbum, spotify.AlbumTypeSingle, spotify.AlbumTypeAppearsOn, spotify.AlbumTypeCompilation}
		result, err := client.GetArtistAlbums(ctx, spotify.ID(artistID), albumTypes, spotify.Market("GB"), spotify.Offset(offset), spotify.Limit(50))
		if err != nil {
			panic(err)
		}
		results = append(results, extractAlbumDetails(result.Albums)...)
		if result.Next == "" {
			complete = true
		}
		offset += 50
	}
	return results
}

func GetArtistAlbumsAfterDate(ctx context.Context, client *spotify.Client, artistID string, date time.Time) []*model.Album {
	albums := GetArtistAlbums(ctx, client, artistID)
	var filteredAlbums []*model.Album
	for _, album := range albums {
		if album.ReleaseDate.After(date) {
			filteredAlbums = append(filteredAlbums, album)
		}
	}
	printAlbumDetails(filteredAlbums)
	return filteredAlbums
}

func extractAlbumDetails(albums []spotify.SimpleAlbum) []*model.Album {
	var extractedAlbums []*model.Album
	for _, album := range albums {
		extractedAlbums = append(extractedAlbums, &model.Album{
			ID:          album.ID.String(),
			Name:        album.Name,
			ReleaseDate: album.ReleaseDateTime(),
		})
	}
	return extractedAlbums
}

func printAlbumDetails(albums []*model.Album) {
	for _, album := range albums {
		fmt.Println(album.Name + " " + album.ReleaseDate.String())
	}
}
