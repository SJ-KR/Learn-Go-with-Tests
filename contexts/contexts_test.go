package contexts

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {

}
func TestHandler(t *testing.T) {
	t.Run("test store", func(t *testing.T) {
		data := "Hello, World"
		server := Server(&StubStore{data})

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s, want %s", response.Body.String(), data)
		}
	})
	t.Run("tells store to cancel work", func(t *testing.T) {
		store := &SpyStore{}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(time.Millisecond*5, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}
	})
	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: "data"}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != "data" {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), "data")
		}
		if store.cancelled {
			t.Errorf("it should not have cancelled the store")
		}
	})
}
