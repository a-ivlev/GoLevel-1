package main

import "fmt"

// fibonacci рекурсивный метод вычисления числа Фибоначчи.
func fibonacci(n uint64) uint64 {
	if n == 2 || n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

// fibonacciClassic классический метод вычисления числа Фибоначчи.
func fibonacciClassic(n uint64) uint64 {
	var f1, f2 uint64
	f2 = 1

	for i := uint64(0); i < n-1; i++ {
		f1, f2 = f2, f1+f2
	}

	return f2
}

func fibonacciClassicCache(n uint64, cache map[uint64]uint64) uint64 {
	var f1, f2 uint64
	f2 = 1

	if _, ok := cache[n]; !ok {
		for i := uint64(0); i < n-1; i++ {
			f1, f2 = f2, f1+f2
			cache[i+2] = f2
		}
	}

	return cache[n]
}

// fibonacciCache Оптимизация приложения за счёт сохранения предыдущих результатов в мапе.
func fibonacciCache(n uint64, cache map[uint64]uint64) uint64 {
	if n == 2 || n == 1 {
		return 1
	}

	if _, ok := cache[n]; !ok {
		cache[n] = fibonacciCache(n-1, cache) + fibonacciCache(n-2, cache)
	}

	return cache[n]
}

func main() {

	n:= uint64(10)

	cache := make(map[uint64]uint64, n)

	fmt.Println("fib = ", fibonacciCache(n, cache))

}

