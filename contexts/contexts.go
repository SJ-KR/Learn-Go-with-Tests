package contexts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

/*
type Store interface {
	Fetch() string
	Cancel()
}
*/

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func (s *StubStore) Cancel() {

}

/*
type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}
*/
type SpyStore struct {
	response string
	t        *testing.T
}
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
func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("Spy store got cancelled")
				return
			default:
				time.Sleep(time.Millisecond * 10)
				result += string(c)
			}
		}
		data <- result
	}()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

/*
func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}
func (s *SpyStore) Cancel() {
	s.cancelled = true
}
*/
/*
func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Errorf("store was not told to cancel")
	}

}
func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Errorf("store was told to cancel")
	}
}
*/
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return
		}
		fmt.Fprint(w, data)
	}
}

/*
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()
		select {
		case d := <-data:
			_, _ = fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
*/
