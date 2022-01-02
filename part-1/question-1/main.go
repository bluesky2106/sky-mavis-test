package main

import (
	"fmt"
	"log"
)

func spiralPrint(N int) {
	// number of rows and columns
	nrows, ncols := N, N
	row, col := 0, 0
	for row < nrows && col < ncols {
		// print the first row
		for i := col; i < ncols; i++ {
			num := N*row + i + 1
			fmt.Printf("%d ", num)
		}
		row += 1

		// print the last column from remaining columns
		for i := row; i < nrows; i++ {
			num := N*i + ncols
			fmt.Printf("%d ", num)
		}
		ncols -= 1

		// print the last row from the remaining rows, from right to left
		for i := ncols - 1; i >= col; i-- {
			num := N*ncols + i + 1
			fmt.Printf("%d ", num)
		}
		nrows -= 1

		// print the first column from the remaining columns from bottom to top
		for i := nrows - 1; i >= row; i-- {
			num := N*i + col + 1
			fmt.Printf("%d ", num)
		}
		col += 1
	}
}

func main() {
	var N int
	fmt.Println("Enter a positive number:")
	_, err := fmt.Scan(&N)
	if err != nil {
		log.Fatal(err)
	}
	if N < 0 {
		log.Fatal("Number must be bigger than 0")
	}
	spiralPrint(N)
}
