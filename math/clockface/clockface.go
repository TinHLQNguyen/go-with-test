package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInFullClock = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInFullClock = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInFullClock   = 2 * hoursInHalfClock
)

// A Point represent a 2D Catersian coordinate
type Point struct {
	X float64
	Y float64
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(t.Second()))
}

func minutesInRadians(t time.Time) float64 {
	return secondsInRadians(t)/minutesInFullClock +
		math.Pi/(minutesInHalfClock/float64(t.Minute()))
}

func hoursInRadians(t time.Time) float64 {
	return minutesInRadians(t)/hoursInFullClock +
		math.Pi/(hoursInHalfClock/float64(t.Hour()%hoursInFullClock))
}

func secondsHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minutesHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hoursHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
