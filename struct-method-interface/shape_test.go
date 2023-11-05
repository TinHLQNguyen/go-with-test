package main

import "testing"

func assertCorrectValue(t testing.TB, when, then float64) {
	t.Helper()
	if when != then {
		t.Errorf("when %.2f then %.2f", when, then)
	}
}

// use given-when-then pattern
func TestPerimeter(t *testing.T) {
	given := Rectangle{10.0, 10.0}
	when := Perimeter(given)
	then := 40.0

	assertCorrectValue(t, when, then)
}

func TestArea(t *testing.T) {
	given := Rectangle{10.0, 10.0}
	when := Area(given)
	then := 100.0

	assertCorrectValue(t, when, then)
}
