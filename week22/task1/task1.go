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

func countElemAfterMatch(array [size]int, val int) int {
	index := len(array) - 1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			break
		}
	}
	return len(array) - index - 1
}

func main() {
    rand.Seed(time.Now().UnixNano())

	array := generateRandomArray()
	fmt.Println("Сгенерированный массив:", array)

	var val int
	fmt.Println("Введите число:")
	fmt.Scan(&val)
	
	fmt.Println("Количество чисел в массиве после введёного:", countElemAfterMatch(array, val))
}