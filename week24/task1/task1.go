package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func generateRandomArray() [size]int {
	var array [size]int
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(size)
	}
	return array
}

func insertionSort(array [size]int) [size]int {
	for i := 1; i < len(array); i++ {
		for j := i - 1; j >= 0 && array[j + 1] < array[j]; j-- {
			array[j+1], array[j] = array[j], array[j+1]
		}
	}
	return array
}

func main() {
    rand.Seed(time.Now().UnixNano())

	array := generateRandomArray()
	fmt.Println("Сгенерированный массив:", array)

	fmt.Println("Отсортированный массив:", insertionSort(array))
}