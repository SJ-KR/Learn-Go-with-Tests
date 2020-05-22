package syncs

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter", func(t *testing.T) {
		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
}
func assertCounter(t *testing.T, got Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), 3)
	}
}
