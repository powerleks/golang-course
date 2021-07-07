package main

import "fmt"

func isEven(number int) bool {
	return number % 2 == 0
}

func main() {
	fmt.Println("Данная программа определет, является число четным")

	fmt.Println(2, isEven(2))
	fmt.Println(3, isEven(3))
}
