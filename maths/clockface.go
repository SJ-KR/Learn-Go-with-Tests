package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const secondHandLength = 90
const CentreX = 150
const CentreY = 150

func SecondHand(t time.Time) Point {
	p := SecondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + CentreX, p.Y + CentreY}
	return p
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
func SecondHandPoint(t time.Time) Point {
	a := SecondsInRadian(t)
	x := math.Sin(a)
	y := math.Cos(a)
	return Point{x, y}
}
