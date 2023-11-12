package main

import (
	"bytes"
	"fmt"
)

func CountDown(out *bytes.Buffer) {
	fmt.Fprint(out, "3")
}

// func main() {
// 	CountDown()
// }
