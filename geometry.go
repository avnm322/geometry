package geometry

import "math"

// Интерфейс для геометрических фигур
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Структура для круга
type Circle struct {
	Radius float64
}

// Метод для вычисления площади круга
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Метод для вычисления периметра круга
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Структура для прямоугольника
type Rectangle struct {
	Width  float64
	Height float64
}

// Метод для вычисления площади прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Метод для вычисления периметра прямоугольника
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}
