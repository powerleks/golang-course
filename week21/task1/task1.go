package main

import (
	"fmt"
)

func calculate(x int16, y uint8, z float32) float32 {
	s := 2 * float32(x) + float32(y) * float32(y) - 3 / z
	return s
}

func main() {
	var x int16
	var y uint8
	var z float32

	fmt.Println("Введите значение x")
	fmt.Scan(&x)
	fmt.Println("Введите значение y")
	fmt.Scan(&y)
	fmt.Println("Введите значение z")
	fmt.Scan(&z)
	
	fmt.Println("S =", calculate(x, y, z))
}