package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Run("tell store to cancel work if request is cancelled", func(t *testing.T) {
		data := "Hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// derive new context (ctx) from original request, get the cancel func along
		cancellingCtx, cancel := context.WithCancel(request.Context())
		// schedule cancel func to be called
		time.AfterFunc(5*time.Millisecond, cancel)
		// use this new context for request
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}
	})
	t.Run("return data from store", func(t *testing.T) {
		data := "Hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		if store.cancelled {
			t.Errorf("store should not have cancelled")
		}
	})
}
