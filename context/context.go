package mycontext

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

///// Stubstore cannot fulfil condition of Store interface anymore, use SpyStore for test instead
// type StubStore struct {
// 	response string
// }

// func (s *StubStore) Fetch() string {
// 	return s.response
// }

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		// running data in goroutine will race Fetch and Cancel signal
		go func() {
			data <- store.Fetch()
		}()

		select {
		// data fetched normally
		case d := <-data:
			fmt.Fprint(w, d)
		// cancel func is called
		case <-ctx.Done():
			store.Cancel()
		}

	}
}