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
	areaTests := []struct {
		shape Shape
		then  float64
	}{
		{Rectangle{10.0, 10.0}, 100.0},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		// define in-loop var to prevent this https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		shape := tt.shape
		then := tt.then
		assertShapeArea(t, shape, then)
	}
}
