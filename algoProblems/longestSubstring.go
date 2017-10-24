package main

import "fmt"

// Given a string, find the length of the longest substring without repeating characters.
// Examples:
// Given "abcabcbb", the answer is "abc", which the length is 3.
// Given "bbbbb", the answer is "b", with the length of 1.
// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

// this is my misinterpretation of the problem
// only repeatitions that count are if we add
// the same character twice in a row
func findLongestSubstring(input string) string {
	characters := []byte(input)
	var finalSubstring []byte
	var currentSubstring []byte

	for index, character := range characters {
		currentSubstring = append(currentSubstring, character)
		// early exit case, if we're on the last character then
		// the current finalSubstring should be the longest substring
		if index == len(characters)-1 {
			break
		}
		// if the next character is the same as the current character
		if characters[index+1] == character {
			if len(currentSubstring) > len(finalSubstring) {
				finalSubstring = currentSubstring
				currentSubstring = nil
				continue
			} else {
				continue
			}
		}
	}
	return string(finalSubstring)
}

func isCharInSlice(characters []byte, character byte) bool {
	for _, char := range characters {
		if character == char {
			return true
		}
	}
	return false
}

// proper solution, the substring can't contain any duplicate characters
func findLongestSubstringNoDuplicates(input string) string {
	characters := []byte(input)
	var finalSubstring []byte
	var currentSubstring []byte

	for _, character := range characters {
		if isCharInSlice(currentSubstring, character) == false {
			currentSubstring = append(currentSubstring, character)
			if len(currentSubstring) > len(finalSubstring) {
				fmt.Println("currentSubstring: ", string(currentSubstring))
				fmt.Println("finalSubstring: ", string(finalSubstring))
				finalSubstring = currentSubstring
			}
		} else {
			// is there a nicer way to do this?
			currentSubstring = nil
			currentSubstring = append(currentSubstring, character)
		}
	}
	return string(finalSubstring)
}

func main() {
	inputString := "pwwkew"
	fmt.Println(findLongestSubstringNoDuplicates(inputString))

}
