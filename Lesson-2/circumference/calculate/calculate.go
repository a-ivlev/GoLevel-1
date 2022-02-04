package calculate

import "math"

func Diametr(area float64) float64 {
	return math.Sqrt(area*4/3.14)
}

func Circumference(area float64) float64 {
	return math.Sqrt(area*4*3.14)
}