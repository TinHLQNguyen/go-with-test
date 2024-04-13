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

func minutesInRadians(t time.Time) float64 {
	return secondsInRadians(t)/60 +
		math.Pi/(30/float64(t.Minute()))
}

func hoursInRadians(t time.Time) float64 {
	return minutesInRadians(t)/12 +
		math.Pi/(6/float64(t.Hour()%12))
}

func secondsHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minutesHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
