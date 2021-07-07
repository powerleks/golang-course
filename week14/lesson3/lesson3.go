package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateInRange(min, max int) int {
	return rand.Intn(max - min) + min
}

func generatePoint() (int, int) {
	return generateInRange(-1000, 1000), generateInRange(-1000, 1000)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i :=0; i < 3; i++ {
		x1, y1 := generatePoint()
		x2 := 2 * x1 + 10
		y2 := -3 * y1 - 5
		fmt.Printf("Точка: (%d, %d)\n", x1, y1)
		fmt.Printf("Точка после преобрзования: (%d, %d)\n", x2, y2)
		fmt.Println("--------------")
	}
}