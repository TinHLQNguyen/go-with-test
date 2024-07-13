package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
			"Loser":  0,
		},
		[]string{},
	}
	server := &PlayerServer{store: &store}

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

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
	}
	server := &PlayerServer{store: &store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/player/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertEqual(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls), 1)
		}
	})
}

func newGetScoreRequest(name string) (request *http.Request) {
	// the request we'll send to test
	request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

// stub for test, following PlayerStore Interface
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, bool) {
	score, ok := s.scores[name]
	return score, ok
}

// spy on POST calls
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}
