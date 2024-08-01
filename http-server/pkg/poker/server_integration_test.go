package poker_test

import (
	"go-with-test/http-server/pkg/poker"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinAndRetrieveThem(t *testing.T) {
	database, cleanDatabase := poker.CreateTempFile(t, "")
	defer cleanDatabase()
	store, err := poker.NewFileSystemPlayerStore(database)
	poker.AssertNoError(t, err)
	server := mustMakePlayerServer(t, store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		// Assert status
		poker.AssertEqual(t, response.Code, http.StatusOK)
		// Assert content
		poker.AssertEqual(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		// Assert status
		poker.AssertEqual(t, response.Code, http.StatusOK)

		got, err := poker.NewLeague(response.Body)
		if err != nil {
			t.Fatalf("unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}
		want := []poker.Player{
			{"Pepper", 3},
		}
		poker.AssertLeague(t, got, want)
	})
}
