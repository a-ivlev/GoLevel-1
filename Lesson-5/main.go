package main

import "fmt"


func fibonacci(n uint64) uint64 {
	if n == 2 || n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

// Оптимизация приложения за счёт сохранения предыдущих результатов в мапе.
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

