package main

import (
	"fmt"
)

func main() {
	fmt.Println("Данная программа производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.")

	var firstArr [4]int
	var secondArr [5]int
	var mergedArr [9]int

	for i, _ := range firstArr {
		fmt.Printf("Введите %v элемент первого массива:\n", i)
		fmt.Scan(&firstArr[i])
	}

	for i, _ := range secondArr {
		fmt.Printf("Введите %v элемент второго массива:\n", i)
		fmt.Scan(&secondArr[i])
	}

	i := 0
	j := 0
	for i < len(firstArr) && j < len(secondArr) {
		if firstArr[i] < secondArr[j] {
			mergedArr[i + j] = firstArr[i]
			i++
		} else {
			mergedArr[i + j] = secondArr[j]
			j++
		}
	}

	for i < len(firstArr) {
		mergedArr[i + j] = firstArr[i]
		i++
	}
	for j < len(secondArr) {
		mergedArr[i + j] = secondArr[j]
		j++
	}
	fmt.Println("Объединенный массив:", mergedArr)
}