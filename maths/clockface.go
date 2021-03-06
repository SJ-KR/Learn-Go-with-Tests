package clockface

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	CentreX          = 150
	CentreY          = 150
)

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
func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadian(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}
func hoursInRadians(t time.Time) float64 {
	return (MinutesInRadians(t) / 12) + (math.Pi / (6 / float64(t.Hour()%12)))
}
func SecondHandPoint(t time.Time) Point {
	a := SecondsInRadian(t)
	return angleToPoint(a)
}
func MinuteHandPoint(t time.Time) Point {
	a := MinutesInRadians(t)
	return angleToPoint(a)
}
func HourHandPoint(t time.Time) Point {
	a := hoursInRadians(t)
	return angleToPoint(a)
}
func angleToPoint(a float64) Point {
	x := math.Sin(a)
	y := math.Cos(a)
	return Point{x, y}
}

type Svg struct {
	XMLName xml.Name `xml:"svg"`

	Xmlns   string `xml:"xmlns,attr"`
	Width   string `xml:"width,attr"`
	Height  string `xml:"height,attr"`
	ViewBox string `xml:"viewBox,attr"`
	Version string `xml:"version,attr"`

	Circle Circle `xml:"circle"`
	Line   []Line `xml:"line"`
}
type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}
type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

func makeHand(p Point, l float64) Point {
	p = Point{p.X * l, p.Y * l}
	p = Point{p.X, -p.Y}
	p = Point{p.X + CentreX, p.Y + CentreY}

	return p
}
func secondHand(w io.Writer, t time.Time) {
	p := makeHand(SecondHandPoint(t), secondHandLength)

	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}
func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(MinuteHandPoint(t), minuteHandLength)

	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}
func hourHand(w io.Writer, t time.Time) {
	p := makeHand(HourHandPoint(t), hourHandLength)

	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}
func SVGWriter(w io.Writer, t time.Time) {

	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)

}
