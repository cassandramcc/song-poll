package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cassandramcc/songpoll/core"
)

type Server struct {
	Poller *core.Poller
}

func enableCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func (s *Server) GetArtists(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)

	if r.Method != http.MethodGet {
		http.Error(w, r.Method+" not allowed on this endpoint", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("Get Artists")
	artists, err := s.Poller.GetArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}

func StartServer(poller *core.Poller) {
	server := &Server{
		Poller: poller,
	}
	http.HandleFunc("/artists", server.GetArtists)
	http.ListenAndServe(":8090", nil)
}
