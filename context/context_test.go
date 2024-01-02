package mycontext

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	// This select waits for the goroutine above
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// new type to make sure in case of cancellation, no data is written to response
// to use if for test, must make it compatble with ResponseWriter Interface https://pkg.go.dev/net/http#ResponseWriter
type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statuscode int) {
	s.written = true
}

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

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Errorf("a response should not have been written")
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
	})
}
