package main

import (
	"fmt"
)

const size = 12

func readArray() [size]int {
	var array [size]int
	for i := 0; i < size; i++ {
		fmt.Printf("Введите элемент массива на позиции %v:\n", i)
		fmt.Scan(&array[i])
	}
	return array
}

func leftBinarySearch(array [size]int, val int) int {
	index := -1
	left := 0
	right := len(array) - 1
	for left <= right {
		middle := left + (right - left) / 2
		if array[middle] == val {
			index = middle
			right = middle - 1
		} else if array[middle] < val {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return index
}

func main() {
	array := readArray()

	var val int
	fmt.Println("Введите число:")
	fmt.Scan(&val)
	
	fmt.Println("Позиция первого вхождения:", leftBinarySearch(array, val))
}