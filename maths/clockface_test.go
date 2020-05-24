package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	times := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(times)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
func TestSecondsInRadian(t *testing.T) {
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)

	want := math.Pi
	got := SecondsInRadian(thirtySeconds)

	if got != want {
		t.Fatalf("want %v radians, but got %v", want, got)
	}
}
