package main

import (
	"fmt"
)

// Алгоритм сортировки вставками
func insertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{9, 5, 3, 6, 8, 7, 1, 2}

	insertionSort(arr)
	
	fmt.Println(arr)

}