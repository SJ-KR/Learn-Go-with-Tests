package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{150, 60}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}
func testName(t time.Time) string {
	return t.Format(t.String())
}

func SecondsInRadian(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}
