package main

import (
	"fmt"
)

// Алгоритм сортировки вставками
func insertion(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return arr
}

func main() {
	arr := []int{9, 5, 3, 6, 8, 7, 1, 2}

	arr = insertion(arr)
	
	fmt.Println(arr)

}