package main

import "fmt"

const size = 3

func calculateDeterminant(matA [size][size]int) int {
	det := matA[0][0] * matA[1][1] * matA[2][2]
	det -= matA[0][0] * matA[1][2] * matA[2][1]
	det -= matA[0][1] * matA[1][0] * matA[2][2]
	det += matA[0][1] * matA[1][2] * matA[2][0]
	det += matA[0][2] * matA[1][0] * matA[2][1]
	det -= matA[0][2] * matA[1][1] * matA[2][0]
	return det
}

func readMatrix() [size][size]int {
	var matrix [size][size]int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("Введите элемент матрицы на позиции (%v, %v):\n", i, j)
			fmt.Scan(&matrix[i][j])
		}
	}
	return matrix
}

func main() {
	fmt.Println("Данная программа вычисляет определитель матрицы 3х3")

	var matrix = readMatrix()
	
	fmt.Println("Определитель:", calculateDeterminant(matrix))
}