package main

import (
	"log"
	"net/http"
)

// kept empty and hard-coded response until one is implemented
type InMemoryPlayerStore struct{}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	return 123, true
}

func (s *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
