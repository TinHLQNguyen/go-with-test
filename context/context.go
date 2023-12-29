package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
}

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}
