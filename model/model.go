package model

import "time"

type Artist struct {
	ID         string
	Name       string
	LastVisted time.Time
}

type Album struct {
	ID          string
	Name        string
	Artists     []Artist
	ReleaseDate time.Time
}

type Track struct {
	ID      string
	Name    string
	Artists []Artist
	Album   Album
}
