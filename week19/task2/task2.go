package main

import (
	"fmt"
)

func main() {
	fmt.Println("Данная программа сортирует пузырьком массив длиной шесть")

	var arr [6]int

	for i, _ := range arr {
		fmt.Printf("Введите %v элемент массива:\n", i)
		fmt.Scan(&arr[i])
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println("Отсортированный массив:", arr)
}
