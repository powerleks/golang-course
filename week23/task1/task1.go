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
		array[i] = rand.Intn(size * 10)
	}
	return array
}

func isEven(val int) bool {
	return val % 2 == 0
}

func breakArrayIntoEvenAndOdd(array [size]int) ([]int, []int) {
	var even = make([]int, 0)
	var odd = make([]int, 0)
	for i := 0; i < len(array); i++ {
		if isEven(array[i]) {
			even = append(even, array[i])
		} else {
			odd = append(odd, array[i])
		}
	}
	return even, odd
}

func main() {
    rand.Seed(time.Now().UnixNano())

	array := generateRandomArray()
	fmt.Println("Сгенерированный массив:", array)

	even, odd := breakArrayIntoEvenAndOdd(array)
	
	fmt.Println("Четные числа в массиве:", even)
	fmt.Println("Нечётные числа в массиве:", odd)
}