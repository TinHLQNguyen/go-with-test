package context

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	data := "Hello, world"
	svr := Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}
