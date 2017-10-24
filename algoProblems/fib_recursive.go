package main

import "fmt"

// 2 minutes 53 seconds, most of the last 53 seconds was missing go imports and error
// find the n-th fibbonacci number
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println("fib 5 is: ", fib(9))
}
