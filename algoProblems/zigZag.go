package main

import "fmt"

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
// Write the code that will take a string and make this conversion given a number of rows:
// string convert(string text, int nRows);
// convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

// notes, the zig zag can be treated like a wave
// use cycles, I drew out what this string would look like with 4 rows

func turnStringToZigZag(input string, rows int) string {
	characters := []byte(input)
	var result []byte

	if rows == 1 { // special case
		return input
	}

	// number of letters per cycle  is 2 * number of rows - 2
	cycleSize := 2*rows - 2

	for i := 0; i < rows; i++ {
		// for each cycle in the string
		for appendIndex := i; appendIndex < len(characters); appendIndex += cycleSize {
			fmt.Println("adding char: ", string(characters[appendIndex]))
			// always append the current character
			result = append(result, characters[appendIndex])
			// determine if another character should be present in
			// this row in this current cycle
			// if so it should be at the index equal to the current
			// appendIndex plus the cycle size minus (row number * 2)
			nextIndex := appendIndex + cycleSize - (i * 2)
			if i != 0 && i != rows-1 && nextIndex < len(characters) {
				fmt.Println("adding char next: ", string(characters[nextIndex]))
				fmt.Println("index is: ", nextIndex)
				result = append(result, characters[nextIndex])
			}
		}
	}
	return string(result)
}

func main() {
	input := "PAYPALISHIRING"
	rows := 4

	fmt.Println(turnStringToZigZag(input, rows))
}
