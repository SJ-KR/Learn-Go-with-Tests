package syncs

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter", func(t *testing.T) {
		counter := NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
	t.Run("safely concurrently", func(t *testing.T) {
		want := 1000
		counter := NewCounter()

		var waitgroup sync.WaitGroup
		waitgroup.Add(want)
		for i := 0; i < want; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&waitgroup)
		}
		waitgroup.Wait()
		assertCounter(t, counter, want)
	})

}
func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), 3)
	}
}
