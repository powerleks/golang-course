package main

import (
	"fmt"
	"math"
)

func deferFunction(a, b int, F func (int, int) int) (resp int) {
	defer func() {
		resp = F(a, b)
	}()
	return
}

func main() {
	fmt.Println("Данная программа перемножает матрицы размером 3х5 и 5х4")

	var a int
	var b int

	fmt.Println("Введите значение a")
	fmt.Scan(&a)
	fmt.Println("Введите значение b")
	fmt.Scan(&b)

	fmt.Printf("%v + %v = %v\n", a, b, deferFunction(a, b, func(x, y int) int {return x + y}))
	fmt.Printf("%v * %v = %v\n", a, b, deferFunction(a, b, func(x, y int) int {return x * y}))
	fmt.Printf("%v ^ %v = %v\n", a, b, deferFunction(a, b, func(x, y int) int {return int(math.Pow(float64(x), float64(y)))}))
}
