package model

import "time"

type Artist struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	LastVisted time.Time `json:"last_visited"`
}

type Album struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Artists     []Artist  `json:"artists"`
	ReleaseDate time.Time `json:"release_date"`
}

type Track struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Artists []Artist `json:"artists"`
	Album   Album    `json:"album"`
}
