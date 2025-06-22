package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cassandramcc/songpoll/src/api"
	"github.com/cassandramcc/songpoll/src/core"
	"github.com/cassandramcc/songpoll/src/model"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/zmb3/spotify/v2"
)

const redirectURI = "http://localhost:8080/callback"

type Server struct {
	Poller        *core.Poller
	SpotifyClient *spotify.Client
	auth          *spotifyauth.Authenticator
}

var (
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

type Login struct {
	OK       bool   `json:"OK"`
	Username string `json:"username"`
}

func enableCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)

	url := s.auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Logged in as:", user.ID)
	s.SpotifyClient = client

	login := Login{
		OK:       true,
		Username: user.ID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(login)
}

func (s *Server) completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := s.auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// use the token to get an authenticated client
	client := spotify.New(s.auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}

func (s *Server) GetArtists(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)

	if r.Method != http.MethodGet {
		http.Error(w, r.Method+" not allowed on this endpoint", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("REQUEST FOR: get artists")
	artists, err := s.Poller.GetArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}

func (s *Server) SeachArtists(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)

	fmt.Println("REQUEST FOR: search artists")

	queryParams := r.URL.Query()

	query := queryParams.Get("query")
	if query == "" {
		http.Error(w, "search query cannot be empty", 500)
		return
	}

	artists, err := api.SearchForArtist(context.Background(), query, s.SpotifyClient)
	if err != nil {
		http.Error(w, "failed to search for artists", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}

func (s *Server) AddArtist(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)

	fmt.Println("REQUEST FOR: add artist")

	if r.Method != http.MethodPost {
		http.Error(w, r.Method+" not allowed on this endpoint", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var artist *model.Artist
	err := json.NewDecoder(r.Body).Decode(&artist)
	if err != nil {
		http.Error(w, "error decoding body", 400)
		return
	}
	artist.LastVisted = time.Now().UTC()
	err = s.Poller.AddArtist(artist)
	if err != nil {
		http.Error(w, "failed to add artist", 500)
	}
}

func StartServer(poller *core.Poller) {

	clientID := os.Getenv("SPOTIFY_ID")
	if clientID == "" {
		panic(errors.New("client ID is empty"))
	}
	clientSecret := os.Getenv("SPOTIFY_SECRET")
	if clientSecret == "" {
		panic(errors.New("client secret is empty"))
	}
	server := &Server{
		Poller: poller,
		auth:   spotifyauth.New(clientID, clientSecret, spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopeUserLibraryRead, spotifyauth.ScopePlaylistModifyPrivate)),
	}

	http.HandleFunc("/login", server.Login)
	http.HandleFunc("/callback", server.completeAuth)
	http.HandleFunc("/artists", server.GetArtists)
	http.HandleFunc("/spotify/artists", server.SeachArtists)
	http.HandleFunc("/artist", server.AddArtist)
	http.ListenAndServe(":8080", nil)
}
