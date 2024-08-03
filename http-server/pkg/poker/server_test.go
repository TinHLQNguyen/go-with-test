package poker_test

import (
	"fmt"
	"go-with-test/http-server/pkg/poker"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestGetPlayers(t *testing.T) {
	store := poker.NewStubPlayerStore(
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
			"Loser":  0,
		},
		[]string{},
		nil)
	server := mustMakePlayerServer(t, store)

	t.Run("Return Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		// mock with a spy built in
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		// assert status code
		poker.AssertEqual(t, response.Code, http.StatusOK)
		// assert response body
		poker.AssertEqual(t, response.Body.String(), "20")
	})

	t.Run("Return Loser's 0 score", func(t *testing.T) {
		request := newGetScoreRequest("Loser")
		// mock with a spy built in
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		// assert status code
		poker.AssertEqual(t, response.Code, http.StatusOK)
		// assert response body
		poker.AssertEqual(t, response.Body.String(), "0")
	})

	t.Run("Return 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("SomeGuy")
		// mock with a spy built in
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertEqual(t, response.Code, http.StatusNotFound)
	})
}

func newGetScoreRequest(name string) (request *http.Request) {
	// the request we'll send to test
	request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func TestStoreWins(t *testing.T) {
	store := poker.NewStubPlayerStore(nil, nil, nil)
	server := mustMakePlayerServer(t, store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		// assert status code
		poker.AssertEqual(t, response.Code, http.StatusAccepted)

		poker.AssertPlayerWin(t, store, player)
	})
}

func TestLeague(t *testing.T) {
	wantedLeague := []poker.Player{
		{"Abe", 10},
		{"Bob", 22},
		{"Cleo", 30},
	}
	store := poker.NewStubPlayerStore(nil, nil, wantedLeague)
	server := mustMakePlayerServer(t, store)

	t.Run("it returns StatusOK 200 on /league", func(t *testing.T) {
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got, err := poker.NewLeague(response.Body)
		if err != nil {
			t.Fatalf("unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}
		// assert status code
		poker.AssertEqual(t, response.Code, http.StatusOK)
		poker.AssertLeague(t, got, wantedLeague)

		assertContentType(t, response, "application/json")
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, &poker.StubPlayerStore{})

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertEqual(t, response.Code, http.StatusOK)
	})
	t.Run("When we get a winner over websocket, it's the winner", func(t *testing.T) {
		store := &poker.StubPlayerStore{}
		winner := "Ruth"
		server := httptest.NewServer(mustMakePlayerServer(t, store))
		defer server.Close()

		wsUrl := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

		ws := mustDialWS(t, wsUrl)
		defer ws.Close()

		writeWSMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond) // FIXME random sleep delay to make sure ws could have time to read and RecordWin
		poker.AssertPlayerWin(t, store, winner)
	})
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	if response.Header().Get("content-type") != want {
		t.Errorf("response did not have valid content-type of %s, got %v", want, response.Result().Header)
	}
}

func newPostWinRequest(name string) (request *http.Request) {
	request, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func newGameRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return req
}

func mustMakePlayerServer(t *testing.T, store poker.PlayerStore) *poker.PlayerServer {
	server, err := poker.NewPlayerServer(store)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}
	return ws
}

func writeWSMessage(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws %v", err)
	}
}
