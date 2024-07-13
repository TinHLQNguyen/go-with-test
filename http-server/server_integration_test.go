package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinAndRetrieveThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := &PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	// Assert status
	AssertEqual(t, response.Code, http.StatusOK)
	// Assert content
	AssertEqual(t, response.Body.String(), "3")
}
