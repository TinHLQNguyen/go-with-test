package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	t.Run("Return Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		// mock with a spy built in
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		AssertEqual(t, response.Body.String(), "20")
	})
	t.Run("Return Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		// mock with a spy built in
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		AssertEqual(t, response.Body.String(), "10")
	})
}

func newGetScoreRequest(name string) (request *http.Request) {
	// the request we'll send to test
	request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}
