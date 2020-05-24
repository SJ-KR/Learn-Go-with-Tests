package clockface_test

import (
	clockface "learn-go/maths"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	times := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(times)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
