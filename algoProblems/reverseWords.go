package main

// started 16:45

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

// reverse words in a string but preserve word order and whitespace
func reverseWords(sentence string) string {
	chunks := strings.Split(sentence, " ")
	var revChunks []string
	for _, word := range chunks {
		revChunks = append(revChunks, revString(word))
	}
	return strings.Join(revChunks, " ")
}

func main() {
	fmt.Println(reverseWords("this is a sentence. with some words."))
}
