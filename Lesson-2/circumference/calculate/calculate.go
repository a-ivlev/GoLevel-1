package calculate

import "math"

// Diametr принимает на вход площадь круга и вычисляет диаметр.
func Diametr(area float64) float64 {
	return math.Sqrt(area*4/3.14)
}

// Circumference принимает на взод площадь круга и вычисляет длину окружности.
func Circumference(area float64) float64 {
	return math.Sqrt(area*4*3.14)
}