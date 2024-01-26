package clockface

import (
	"math"
	"time"
)

// A Point represent a 2D Catersian coordinate
type Point struct {
	X float64
	Y float64
}

const secondHandLength = 90
const clockCenterX = 150
const clockCenterY = 150

// SecondHand is the unit vector of the second hand of an analogue clock at time `t` represented as a Point.
func SecondHand(t time.Time) Point {
	p := secondsHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scaling
	p = Point{p.X, -p.Y}                                      // flip coordinate
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}         // translate coordinate
	return p
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func secondsHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
