package main

import "testing"

func assertCorrectValue(t testing.TB, when, then float64) {
	t.Helper()
	if when != then {
		t.Errorf("when %g then %g", when, then)
	}
}

// use given-when-then pattern
func TestPerimeter(t *testing.T) {
	t.Run("perimeter of a rectangle", func(t *testing.T) {
		rectange := Rectangle{10.0, 10.0}
		when := rectange.Perimeter()
		then := 40.0

		assertCorrectValue(t, when, then)
	})
	t.Run("perimeter of a circle", func(t *testing.T) {
		circle := Circle{5.0}
		when := circle.Perimeter()
		then := 31.41592653589793

		assertCorrectValue(t, when, then)
	})
}

func TestArea(t *testing.T) {
	t.Run("area of a rectangle", func(t *testing.T) {
		// given
		rectange := Rectangle{10.0, 10.0}
		when := rectange.Area()
		then := 100.0

		assertCorrectValue(t, when, then)
	})
	t.Run("area of a circle", func(t *testing.T) {
		// given
		circle := Circle{10.0}
		when := circle.Area()
		then := 314.1592653589793

		assertCorrectValue(t, when, then)
	})
}
