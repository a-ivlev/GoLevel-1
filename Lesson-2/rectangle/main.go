package main

import "fmt"


func main() {
	var a, b float64

	fmt.Println("Введите чему равна сторона а прямоугольника: ")
	fmt.Scanln(&a)

	fmt.Println("Введите чему равна сторона b прямоугольника: ")
	fmt.Scanln(&b)

	fmt.Printf("Площадь прямоугодьника равна: %.2f\n", a * b)
}