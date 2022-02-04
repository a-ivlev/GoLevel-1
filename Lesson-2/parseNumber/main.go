package main

import (
	"fmt"

	"github.com/GoLevel-1/errApp"

	"go.uber.org/zap"
)


// userInput просит пользователя ввести данные и проверяет данные на валидность.
func userInput() (int, error) {
	var s int

	fmt.Println("Введите трёхзначное число от 100 до 999: ")
	fmt.Scanln(&s)

	switch {
	case s < 100:
		return 0, errApp.ErrNumberLess100
	case s > 999:
		return 0, errApp.ErrNumberGreater999
	default:
	}

	return s, nil
}

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	num, err := userInput()
	if err != nil {
		logger.Error(fmt.Sprintf(`user input error: %s`, err))
		return
	}

	str := fmt.Sprint(num)

	fmt.Println("Количество сотен ", string(str[0]))
	fmt.Println("Количество десятков ", string(str[1]))
	fmt.Println("Количество единиц ", string(str[2]))

	edin := num % 10
	des := ((num-edin)% 100)/10
	sot := ((num - edin - des)%1000)/100

	fmt.Println()
	fmt.Println("Количество сотен ", sot)
	fmt.Println("Количество десятков ", des)
	fmt.Println("Количество единиц ", edin)
}