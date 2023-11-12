package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySlepper := &SpySleeper{}

	CountDown(buffer, spySlepper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySlepper.Calls != 3 {
		t.Errorf("not enough calls to sleep")
	}
}
