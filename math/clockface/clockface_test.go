package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
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

			want := c.angle
			got := secondsInRadians(c.time)

			if got != want {
				t.Errorf("Got %v, want %v radians", got, want)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), (math.Pi / (30 * 60)) * 7},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {

			want := c.angle
			got := minutesInRadians(c.time)

			if got != want {
				t.Errorf("Got %v, want %v radians", got, want)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {

			want := c.point
			got := secondsHandPoint(c.time)

			if !roughlyEqualPoint(got, want) {
				t.Errorf("Got %v, want %v Point", got, want)
			}

		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	jpLocation := time.FixedZone("JST", +9*int(time.Hour/time.Second))
	return time.Date(2024, time.January, 1, hours, minutes, seconds, 0, jpLocation)
}

func testName(t time.Time) string {
	return t.Format("10:00:05")
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}
