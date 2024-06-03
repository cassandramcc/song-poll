package model

import "time"

type Artist struct {
	ID         string    `csv:"uri"`
	Name       string    `csv:"name"`
	LastVisted time.Time `csv:"last_visited"`
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
