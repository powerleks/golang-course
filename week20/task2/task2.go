package main

import "fmt"

const (
	rowsA = 3
	colsA = 5
	rowsB = 5
	colsB = 4
)

func scalarProduct(matA [rowsA][colsA]int, matB [rowsB][colsB]int, row, col int) int {
	product := 0
	for i := 0; i < colsA; i++ {
		product += matA[row][i] * matB[i][col]
	}
	return product
}

func matrixMult(matA [rowsA][colsA]int, matB [rowsB][colsB]int) [rowsA][colsB]int {
	var matC [rowsA][colsB]int
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			matC[i][j] = scalarProduct(matA, matB, i, j)
		}
	}
	return matC
}

func readMatrixA() [rowsA][colsA]int {
	var matrix [rowsA][colsA]int

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			fmt.Printf("Введите элемент первой матрицы на позиции (%v, %v):\n", i, j)
			fmt.Scan(&matrix[i][j])
		}
	}

	return matrix
}

func readMatrixB() [rowsB][colsB]int {
	var matrix [rowsB][colsB]int

	for i := 0; i < rowsB; i++ {
		for j := 0; j < colsB; j++ {
			fmt.Printf("Введите элемент второй матрицы на позиции (%v, %v):\n", i, j)
			fmt.Scan(&matrix[i][j])
		}
	}
	
	return matrix
}

func main() {
	fmt.Println("Данная программа перемножает матрицы размером 3х5 и 5х4")

	var matrixA = readMatrixA()

	var matrixB = readMatrixB()

	matrix := matrixMult(matrixA, matrixB)

	fmt.Println("Результат перемножения матриц:", matrix)
}
