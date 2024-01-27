package main

import (
	"fmt"
	"go-with-test/math/clockface"
)

func main() {
	// t := time.Now()
	// sh := clockface.SecondHand(t)
	// io.WriteString(os.Stdout, svgStart)
	// io.WriteString(os.Stdout, bezel)
	// io.WriteString(os.Stdout, secondHandTag(sh))
	// io.WriteString(os.Stdout, svgEnd)
}

func secondHandTag(p clockface.Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}
