package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("Площадь прямоугольника: %.2f\n", rect.Area())
}