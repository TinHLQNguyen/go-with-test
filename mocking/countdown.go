package main

import (
	"fmt"
	"io"
	"os"
)

func CountDown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, "Go!")
}

func main() {
	CountDown(os.Stdout)
}
