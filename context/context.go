package mycontext

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		// now Sever doesn't handle cancellation anymore.
		// it only pass through context
		// the recepient (Store) needs to handle it
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprint(w, data)
	}
}
