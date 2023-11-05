package main

import "testing"

// use given-when-then pattern
func TestPerimeter(t *testing.T) {
	given := Rectangle{10.0, 10.0}
	when := Perimeter(given)
	then := 40.0

	if when != then {
		t.Errorf("when %.2f then %.2f", when, then)
	}
}

func TestArea(t *testing.T) {
	given := Rectangle{10.0, 10.0}
	when := Area(given)
	then := 100.0

	if when != then {
		t.Errorf("when %.2f then %.2f", when, then)
	}
}
