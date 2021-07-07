package main

import (
	"fmt"
)

func bubbleSort(array ...int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array) - 1 - i; j++ {
			if array[j + 1] > array[j] {
				array[j + 1], array[j] = array[j], array[j + 1]
			}
		}
	}
	return array
}

func main() {

	fmt.Println("Отсортированный массив:", bubbleSort(1, 2, 3, 4, 5, 6, 7, 8, 9))
}