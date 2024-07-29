package poker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// TODO allow concurrency POST & GET with mutex
// TODO implement a real database playerstorage

const jsonContentType = "application/json"

type PlayerStore interface {
	GetPlayerScore(string) (int, bool)
	RecordWin(string)
	GetLeague() League
}

type Player struct {
	Name string
	Wins int
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// generate new server along with router from store DB
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/game", http.HandlerFunc(p.gameHandler))
	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.Header().Set("content-type", jsonContentType)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	w.WriteHeader(http.StatusAccepted)
	p.store.RecordWin(player)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score, ok := p.store.GetPlayerScore(player)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) gameHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
