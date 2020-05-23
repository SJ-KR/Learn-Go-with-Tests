package contexts

import (
	"net/http"
	"net/http/httptest"
	"testing"
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
}
