package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// technically here, we use another type (buffer) which share the same interface as io
	buffer := bytes.Buffer{}
	Greet(&buffer, "Tin")

	got := buffer.String()
	want := "Hello, Tin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
