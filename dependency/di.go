package main

import (
	"fmt"
	"io"
	"os"
)

// this works because io.Writer is the interface of both os.Stdout and bytes.Buffer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Tin")
}
