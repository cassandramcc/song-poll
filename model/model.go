package model

import "time"

type Artist struct {
	URI        string    `json,bson,csv:"uri"`
	Name       string    `json,bson,csv:"name"`
	LastVisted time.Time `json,bson,csv:"last_visited"`
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
