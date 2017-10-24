package main

import "fmt"

// 7 minutes 1 second, at least 1 minute figuring out for loops since I have no internet for docs and 1 minute debugging my off by one error (n + 1)
// find the n-th fibbonacci number
func fib(n int) int {
	var current = 0
	var prev = 0
	var sec_prev = 0
	for i := 1; i < n+1; i++ {
		if i == 1 || i == 2 {
			current = 1
			prev = 1
			sec_prev = 1
		} else {
			current = prev + sec_prev
			sec_prev = prev
			prev = current
		}
	}
	return current
}

func main() {
	fmt.Println("fib 9 is: ", fib(9))
}
