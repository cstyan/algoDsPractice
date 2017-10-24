package main

// reverse a string

import (
	"fmt"
	"strings"
)

func revString(str string) string {
	var rev string
	chars := strings.Split(str, "")
	for i := len(str) - 1; i >= 0; i-- {
		rev = rev + chars[i]
	}
	return rev
}

func revStringInPlace(str *string) {
	chars := strings.Split(*str, "")
	for i, j := 0, len(*str)-1; i < j; i, j = i+1, j-1 {
		temp := chars[i]
		chars[i] = chars[j]
		chars[j] = temp
	}
	*str = strings.Join(chars, "")
}

func main() {
	hello := "hello"
	fmt.Println("rev: ", revString(hello))
	revStringInPlace(&hello)
	fmt.Println("rev 2:", hello)
}
