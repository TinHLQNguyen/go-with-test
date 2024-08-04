package poker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/gorilla/websocket"
)

// TODO: allow concurrency POST & GET with mutex
// TODO: implement a real database playerstorage

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
	tempate *template.Template
	game    Game
}

const htmlTemplatePath = "game.html"

// generate new server along with router from store DB
func NewPlayerServer(store PlayerStore, game Game) (*PlayerServer, error) {
	p := new(PlayerServer)
	tmpl, err := template.ParseFiles(htmlTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("error loading template from %v %s", htmlTemplatePath, err)
	}

	p.game = game
	p.tempate = tmpl
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/game", http.HandlerFunc(p.gameHandler))
	router.Handle("/ws", http.HandlerFunc(p.webSocket))
	p.Handler = router

	return p, nil
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
	p.tempate.Execute(w, nil)
}

func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
	ws := newPlayerServerWS(w, r)

	numOfPlayerMsg := ws.WaitForMsg()
	numOfPlayer, _ := strconv.Atoi(string(numOfPlayerMsg))
	p.game.Start(numOfPlayer, ws)

	winner := ws.WaitForMsg()
	p.game.Finish(string(winner))
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// b/c what's needed is only a portion of websocket, we encapsulate what we need
type playerServerWS struct {
	*websocket.Conn
}

func newPlayerServerWS(w http.ResponseWriter, r *http.Request) *playerServerWS {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("problem upgrading connection to websocket %v\n", err)
	}
	return &playerServerWS{conn}
}

func (w *playerServerWS) WaitForMsg() string {
	_, msg, err := w.ReadMessage()
	if err != nil {
		log.Printf("error reading from websocket %v\n", err)
	}
	return string(msg)
}

// with Write(p []byte) (n int, err error), ws implemented io.Writer interface
func (w *playerServerWS) Write(p []byte) (n int, err error) {
	err = w.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
