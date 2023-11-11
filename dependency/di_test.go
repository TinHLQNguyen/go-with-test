package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer = bytes.Buffer{}
	Greet(&buffer, "Tin")

	got := buffer.String()
	want := "Hello, Tin"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
