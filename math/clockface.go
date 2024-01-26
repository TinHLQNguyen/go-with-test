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

// SecondHand is the unit vector of the second hand of an analogue clock at time `t` represented as a Point.
func SecondHand(t time.Time) Point {
	p := secondsHandPoint(t)
	p = Point{p.X * 90, p.Y * 90}   // scaling
	p = Point{p.X, -p.Y}            // flip coordinate
	p = Point{p.X + 150, p.Y + 150} // translate coordinate
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
