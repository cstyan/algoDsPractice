package main

import (
	"fmt"
)

// given array of numbers, all appear twice except one, find that one

func findNum(arr []int) int {
	vals := make(map[int]int)
	for _, val := range arr {
		vals[val] = vals[val] + 1
	}
	// there should only be one here
	for k, v := range vals {
		if v == 1 {
			return k
		}
	}
	return -1
}

func findNumSingleLoop(arr []int) int {
	vals := make(map[int]int)
	for _, val := range arr {
		if _, ok := vals[val]; ok {
			delete(vals, val)
			continue
		}
		vals[val] = 1
	}
	// we have to use a loop, but there is only one key in the map at this point
	for k, _ := range vals {
		return k
	}
	return -1
}

func main() {
	fmt.Println("num: ", findNum([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15}))
	fmt.Println("num: ", findNumSingleLoop([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15}))
}
