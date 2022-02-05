package main

import (
	"fmt"
	"go.uber.org/zap"
)

func isprime(n int) bool {
	if n == 1 {
		return false // 1 - не простое число
	}
		
	// перебираем возможные делители от 2 до sqrt(n)
	var d int
	for d = 2; d*d <= n; d++{ 
		// если разделилось нацело, то составное
		if n % d == 0 {
			return false
		}  	
	}
	// если нет нетривиальных делителей, то простое
	return true
}

func main() {
	var end int
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()


	fmt.Println("До какого числа необходимо найти простые числа")
	_, err := fmt.Scan(&end)
	if err != nil {
		logger.Error(fmt.Sprint(`user input error: `, err))
		return
	}

	natur := []int{}

	for i := 2; i <= end; i++ {
		if isprime(i) {
			natur = append(natur, i)
		}
	}

	fmt.Println(natur)

}