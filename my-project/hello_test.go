package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tin")
	want := "Hello, Tin"

	if got != want {
		t.Errorf("go %q , want %q", got, want)
	}
}
