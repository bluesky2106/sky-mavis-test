package main

import "fmt"

func calculatePI(arr []int, exceptIndex int) int {
	PI := 1
	for idx, num := range arr {
		if idx == exceptIndex {
			continue
		}
		PI *= num
	}
	return PI
}

func main() {
	arr := []int{2, 1, 3, 4, 5}
	for idx := range arr {
		pi := calculatePI(arr, idx)
		fmt.Printf("%d ", pi)
	}
}
