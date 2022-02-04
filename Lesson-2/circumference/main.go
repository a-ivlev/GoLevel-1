package main

import (
	"fmt"

	"github.com/GoLevel-1/calculate"
	"github.com/GoLevel-1/errApp"

	"go.uber.org/zap"
)


func userInput() (float64, error) {
	var s float64

	fmt.Println("Введите чему равна площадь круга: ")
	fmt.Scanln(&s)

	switch {
	case s < 0:
		return 0, errApp.ErrAreaNegativNumber
	case s == 0:
		return 0, errApp.ErrAreaZero
	default:
	}

	return s, nil
}

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	s, err := userInput()
	if err != nil {
		logger.Error(fmt.Sprintf(`user input error: %s`, err))
		return
	}

	fmt.Printf("Диаметр окружности: %.2f\n", calculate.Diametr(s))
	fmt.Printf("Длина окружности: %.2f\n", calculate.Circumference(s))
}