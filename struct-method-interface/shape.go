package main

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectange Rectangle) float64 {
	return 2 * (rectange.Width + rectange.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
