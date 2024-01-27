package clockface

import (
	"bytes"
	"encoding/xml"
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

func TestSGVWriterAtMidnight(t *testing.T) {
	tm := simpleTime(0, 0, 0)

	b := bytes.Buffer{}
	SVGWriter(&b, tm)

	svg := SVG{}
	xml.Unmarshal(b.Bytes(), &svg)

	x2 := "150"
	y2 := "60"

	for _, line := range svg.Line {
		if line.X2 == x2 && line.Y2 == y2 {
			return
		}
	}

	t.Errorf("Expected to find the second hand with x2 %+v and y2 %+v, in the SVG output %v", x2, y2, b.String())
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

// SVG was generated 2024-01-27 15:41:29 by https://xml-to-go.github.io/ in Ukraine.
// Struct for XML package to unmarshall stuff
type SVG struct {
	// the 3rd element of struct field is tag, used by packages that uses reflect
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  struct {
		Text  string `xml:",chardata"`
		Cx    string `xml:"cx,attr"`
		Cy    string `xml:"cy,attr"`
		R     string `xml:"r,attr"`
		Style string `xml:"style,attr"`
	} `xml:"circle"`
	Line []struct {
		Text  string `xml:",chardata"`
		X1    string `xml:"x1,attr"`
		Y1    string `xml:"y1,attr"`
		X2    string `xml:"x2,attr"`
		Y2    string `xml:"y2,attr"`
		Style string `xml:"style,attr"`
	} `xml:"line"`
}
