package clockface

import (
	"bytes"
	"encoding/xml"
	"testing"
)

func TestSGVWriterAtMidnight(t *testing.T) {
	tm := simpleTime(0, 0, 0)

	b := bytes.Buffer{}
	SVGWriter(&b, tm)

	svg := SVG{}
	xml.Unmarshal(b.Bytes(), &svg)

	x2 := "150.000"
	y2 := "60.000"

	for _, line := range svg.Line {
		if line.X2 == x2 && line.Y2 == y2 {
			return
		}
	}

	t.Errorf("Expected to find the second hand with x2 %+v and y2 %+v, in the SVG output %v", x2, y2, b.String())
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
