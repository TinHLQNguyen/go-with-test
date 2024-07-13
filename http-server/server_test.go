package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	t.Run("Return Pepper's score", func(t *testing.T) {
		// the request we'll send to test
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		// mock with a spy built in
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		AssertEqual(t, got, want)
	})
}
