package clockface

import (
	"bytes"
	"encoding/xml"
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
func TestSecondHandAt30Seconds(t *testing.T) {
	times := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := Point{X: 150, Y: 150 + 90}
	got := SecondHand(times)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
func TestSecondsInRadian(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := SecondsInRadian(c.time)

			if got != c.angle {
				t.Fatalf("want %v radians, but got %v", c.angle, got)
			}
		})

	}
}

func TestSecondHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := SecondHandPoint(c.time)

			if !roughlyEqualPoint(c.point, got) {
				t.Fatalf("want %v radians, but got %v", c.point, got)
			}
		})

	}
}
func TestSVGWriterAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	b := bytes.Buffer{}
	SVGWriter(&b, tm)

	svg := Svg{}
	xml.Unmarshal(b.Bytes(), &svg)

	x2 := "150.000"
	y2 := "60.000"

	for _, line := range svg.Line {
		if line.X2 == x2 && line.Y2 == y2 {
			return
		}
	}
	t.Errorf("Expected to find the second hand with x2 of %+v and y2 of %+v, in the SVG output %v", x2, y2, b.String())
}
func roughlyEqualFloat64(a, b float64) bool {
	const Threshold = 1e-7
	return math.Abs(a-b) < Threshold
}
func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}
