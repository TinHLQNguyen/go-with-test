package main

import (
	"go-with-test/math/clockface/svg"
	"os"
	"time"
)

func main() {
	t := time.Now()
	svg.SVGWriter(os.Stdout, t)
}
