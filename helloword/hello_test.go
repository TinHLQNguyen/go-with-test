package main

import "testing"

// use given-when-then pattern
func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		when := Hello("Tin", "")
		then := "Hello, Tin"

		assertCorrectMessage(t, when, then)
	})

	t.Run("Saying 'Hello, World' when empty input", func(t *testing.T) {
		when := Hello("", "")
		then := "Hello, World"

		assertCorrectMessage(t, when, then)

	})

	t.Run("In Spanish", func(t *testing.T) {
		when := Hello("Tin", "Spanish")
		then := "Hola, Tin"

		assertCorrectMessage(t, when, then)
	})

	t.Run("In French", func(t *testing.T) {
		when := Hello("Tin", "French")
		then := "Bonjour, Tin"

		assertCorrectMessage(t, when, then)
	})
}

func assertCorrectMessage(t testing.TB, when, then string) {
	t.Helper()
	if when != then {
		t.Errorf("go %q , then %q", when, then)
	}
}
