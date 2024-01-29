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

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func secondsHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
