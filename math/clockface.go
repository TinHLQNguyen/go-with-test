package clockface

import "time"

// A Point represent a 2D Catersian coordinate
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t` represented as a Point.
func SecondHand(t time.Time) Point {
	return Point{}
}
