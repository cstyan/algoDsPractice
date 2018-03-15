package main

import "fmt"

// 2 minutes 53 seconds, most of the last 53 seconds was missing go imports and error
// find the n-th fibbonacci number
func fib(n int, fibNums map[int]int) int {
	res := -1
	if val, ok := fibNums[n]; ok {
		return val
	}

	// normal logic
	if n == 0 {
		res = 0
	} else if n == 1 {
		res = 1
	} else {
		res = fib(n-1, fibNums) + fib(n-2, fibNums)
	}
	fibNums[n] = res
	return res
}

func main() {
	n := 101
	// it's probably more efficient to use an array but since we get 0's instead of nils when
	// initializing an array/slice lets use a map
	fibNums := make(map[int]int)
	fmt.Println("fib 101 is: ", fib(n, fibNums))
}
