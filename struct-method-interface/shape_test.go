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
	assertShapePerimeter := func(t testing.TB, shape Shape, then float64) {
		t.Helper()
		when := shape.Perimeter()
		assertCorrectValue(t, when, then)
	}
	t.Run("perimeter of a rectangle", func(t *testing.T) {
		rectange := Rectangle{10.0, 10.0}
		then := 40.0

		assertShapePerimeter(t, rectange, then)
	})
	t.Run("perimeter of a circle", func(t *testing.T) {
		circle := Circle{5.0}
		then := 31.41592653589793

		assertShapePerimeter(t, circle, then)
	})
}

func TestArea(t *testing.T) {
	assertShapeArea := func(t testing.TB, shape Shape, then float64) {
		t.Helper()
		when := shape.Area()
		assertCorrectValue(t, when, then)
	}
	t.Run("area of a rectangle", func(t *testing.T) {
		// given
		rectange := Rectangle{10.0, 10.0}
		then := 100.0

		assertShapeArea(t, rectange, then)
	})
	t.Run("area of a circle", func(t *testing.T) {
		// given
		circle := Circle{10.0}
		then := 314.1592653589793

		assertShapeArea(t, circle, then)
	})
}
