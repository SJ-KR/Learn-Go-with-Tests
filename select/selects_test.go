package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20)
	fastServer := makeDelayedServer(0)

	slowURL := slowServer.URL
	fastURL := fastServer.URL
	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
}
