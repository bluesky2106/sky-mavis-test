package main

import "fmt"

func setToOne(matrix [][]int, rows, cols map[int]bool) {
	for row := range rows {
		vector := matrix[row]
		for col := range vector {
			vector[col] = 1
		}
	}
	for col := range cols {
		for row := range matrix {
			vector := matrix[row]
			vector[col] = 1
		}
	}
}

func modifyMatrix(matrix [][]int) {
	rows := make(map[int]bool, 0)
	cols := make(map[int]bool, 0)
	for row := range matrix {
		vector := matrix[row]
		for col := range vector {
			if vector[col] == 1 {
				rows[row] = true
				cols[col] = true
			}
		}
	}
	setToOne(matrix, rows, cols)
}

func printMatrix2D(matrix [][]int) {
	for row := range matrix {
		vector := matrix[row]
		for col := range vector {
			fmt.Printf("%d ", vector[col])
		}
		fmt.Println()
	}
}

func main() {
	matrix := [][]int{
		{0, 1, 2, 3},
		{3, 1, 2, 4},
		{1, 0, 2, 3},
		{5, 9, 2, 5},
	}
	modifyMatrix(matrix)
	printMatrix2D(matrix)
}
