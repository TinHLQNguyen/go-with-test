package main

import "testing"

// use given-when-then pattern
func TestPerimeter(t *testing.T) {
	when := Perimeter(10.0, 10.0)
	then := 40.0

	if when != then {
		t.Errorf("when %.2f then %.2f", when, then)
	}
}

func TestArea(t *testing.T) {
	when := Area(10.0, 10.0)
	then := 100.0

	if when != then {
		t.Errorf("when %.2f then %.2f", when, then)
	}
}
