package main

import (
	"fmt"
	"math/rand"
	"time"
)

func multiply(n int) (res int) {
	res = n * 100
	return res
}

func add(n int) (res int) {
	res = n + 5
	return res
}

func transform(n int) (res int) {
	value := multiply(n)
	res = add(value)
	return res
}

func main() {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(500)
	res := transform(n)
	fmt.Println("Число до преобразования:", n)
	fmt.Println("Число после преобразования:", res)
}