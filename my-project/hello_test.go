package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Tin")
		want := "Hello, Tin"

		if got != want {
			t.Errorf("go %q , want %q", got, want)
		}
	})

	t.Run("Saying 'Hello, World' when empty input", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("go %q , want %q", got, want)
		}
	})
}
