package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// technically here, we use another type in bytes.Buffer which share the same interface io.Writer as os.Stdout
	buffer := bytes.Buffer{}
	Greet(&buffer, "Tin")

	got := buffer.String()
	want := "Hello, Tin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
