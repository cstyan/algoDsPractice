package main

import (
	"fmt"
	"sort"
)

// find median of two arrays

func findMedian(arr []int) float64 {
	if len(arr)%2 == 0 {
		return float64(arr[len(arr)/2 - 1] + arr[len(arr)/2]) / float64(2)
	}
	return float64(arr[len(arr)/2])
}

func medianTwo(arrOne []int, arrTwo []int) float64 {
	// first, combine both arrays and sort
	result := append(arrOne, arrTwo...)
	sort.Ints(result)
	fmt.Println("result: ", result)
	return findMedian(result)
}

func main() {
	arrOne := [...]int{3, 2, 1, 6, 0, 5}
	arrTwo := [...]int{8, 11, 9, 4, 7}

	fmt.Println("median: ", medianTwo(arrOne[:], arrTwo[:]))
}