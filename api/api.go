package api

import (
	"context"
	"fmt"
	"time"

	"github.com/cassandramcc/songpoll/model"
	"github.com/zmb3/spotify/v2"
)

// GetArtistAlbums retrieves all albums for a given artist
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

// GetArtistAlbumsAfterDate retrieves all albums for a given artist after a given date
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

// GetAlbumTracks retrieves all tracks for a given album
func GetAlbumTracks(ctx context.Context, client *spotify.Client, albumID string) []*model.Track {
	var results []*model.Track
	offset := 0
	complete := false
	for !complete {
		result, err := client.GetAlbumTracks(ctx, spotify.ID(albumID), spotify.Offset(offset), spotify.Limit(50))
		if err != nil {
			panic(err)
		}
		results = append(results, extractTrackDetails(result.Tracks)...)
		if result.Next == "" {
			complete = true
		}
		offset += 50
	}
	printTrackDetails(results)
	return results
}

// Converts a slice of spotify.SimpleAlbums to a slice of model.Albums
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

// Converts a slice of spotify.SimpleTracks to a slice of model.Tracks
func extractTrackDetails(tracks []spotify.SimpleTrack) []*model.Track {
	var extractedTracks []*model.Track
	for _, track := range tracks {
		var artists []model.Artist
		for _, artist := range track.Artists {
			artists = append(artists, model.Artist{
				ID:   artist.ID.String(),
				Name: artist.Name,
			})
		}
		extractedTracks = append(extractedTracks, &model.Track{
			ID:      track.ID.String(),
			Name:    track.Name,
			Artists: artists,
			Album: model.Album{
				ID:   track.Album.ID.String(),
				Name: track.Album.Name,
			},
		})
	}
	return extractedTracks
}

// GetTrackIDs gets the track ids from a slice of tracks
func GetTrackIDs(tracks []*model.Track) []spotify.ID {
	var ids []spotify.ID
	for _, track := range tracks {
		ids = append(ids, spotify.ID(track.ID))
	}
	return ids
}

func printAlbumDetails(albums []*model.Album) {
	for _, album := range albums {
		fmt.Println(album.Name + " " + album.ReleaseDate.String())
	}
}

func printTrackDetails(tracks []*model.Track) {
	for _, track := range tracks {
		fmt.Println(track.Name)
	}
}
