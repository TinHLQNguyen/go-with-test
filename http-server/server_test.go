package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	store := newStubPlayerStore(
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
			"Loser":  0,
		},
		[]string{},
		nil)
	server := NewPlayerServer(store)

	t.Run("Return Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		// mock with a spy built in
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		// assert status code
		AssertEqual(t, response.Code, http.StatusOK)
		// assert response body
		AssertEqual(t, response.Body.String(), "20")
	})

	t.Run("Return Loser's 0 score", func(t *testing.T) {
		request := newGetScoreRequest("Loser")
		// mock with a spy built in
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		// assert status code
		AssertEqual(t, response.Code, http.StatusOK)
		// assert response body
		AssertEqual(t, response.Body.String(), "0")
	})

	t.Run("Return 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("SomeGuy")
		// mock with a spy built in
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertEqual(t, response.Code, http.StatusNotFound)
	})
}

func newGetScoreRequest(name string) (request *http.Request) {
	// the request we'll send to test
	request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func TestStoreWins(t *testing.T) {
	store := newStubPlayerStore(nil, nil, nil)
	server := NewPlayerServer(store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		// assert status code
		AssertEqual(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls), 1)
		}

		// assert if correct name is recorded
		AssertEqual(t, store.winCalls[0], player)
	})
}

func newPostWinRequest(name string) (request *http.Request) {
	request, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Abe", 10},
		{"Bob", 22},
		{"Cleo", 30},
	}
	store := newStubPlayerStore(nil, nil, wantedLeague)
	server := NewPlayerServer(store)

	t.Run("it returns StatusOK 200 on /league", func(t *testing.T) {
		request := newLeaguaRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got, err := NewLeague(response.Body)
		if err != nil {
			t.Fatalf("unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}
		// assert status code
		AssertEqual(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)

		assertContentType(t, response, "application/json")
	})
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	if response.Header().Get("content-type") != want {
		t.Errorf("response did not have valid content-type of %s, got %v", want, response.Result().Header)
	}
}

func newLeaguaRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func assertLeague(t testing.TB, got, want []Player) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GOT %+v, WANT %+v", got, want)
	}
}

// stub for test, following PlayerStore Interface
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	leaque   []Player
}

func newStubPlayerStore(score map[string]int, winCalls []string, leaqueTable []Player) *StubPlayerStore {
	return &StubPlayerStore{score, winCalls, leaqueTable}
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, bool) {
	score, ok := s.scores[name]
	return score, ok
}

// spy on POST calls
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.leaque
}
