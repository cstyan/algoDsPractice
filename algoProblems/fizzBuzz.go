package main

import (
	"fmt"
)

// fizz buzz div 3 == fizz div 5 == buzz div 3 and 5 == buzz

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		divThree := i % 3
		divFive := i % 5
		if divThree == 0 && divFive == 5 {
			fmt.Println("fizzbuzz")
			continue
		} else if divThree == 0 {
			fmt.Println("fizz")
			continue
		} else if divFive == 0 {
			fmt.Println("buzz")
			continue
		}
		fmt.Println(i)
	}
}

func main() {
	fizzBuzz(100)
}
