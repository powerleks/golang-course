package main

import "fmt"

func isEven(number int) bool {
	return number%2 == 0
}

func main() {
	fmt.Println("Данная программа выводит количество четных и нечетных чисел в массиве")

	var array [10]int

	for i, _ := range array {
		fmt.Printf("Введите %v элемент массива:\n", i)
		fmt.Scan(&array[i])
	}

	evenNumberAmount := 0
	oddNumberAmount := 0
	for _, el := range array {
		if el % 2 == 0 {
			evenNumberAmount++
		} else {
			oddNumberAmount++
		}
	}
	fmt.Printf("Количество четных чисел в массиве: %v\n", evenNumberAmount)
	fmt.Printf("Количество нечетных чисел в массиве: %v\n", oddNumberAmount)

}
