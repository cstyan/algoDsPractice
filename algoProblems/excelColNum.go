package main

import (
	"fmt"
	"math"
	"strings"
)

// turn each excel column char into a number
// A = 1, B = 2, ..., AA = 27, AB = 28
// BA is 53, BZ is 78, and CA is 79

// assume we get all valid input, string of all capital letters
func excelNum(excelCol string) int {
	num := 0
	var removal = 64
	chars := strings.Split(excelCol, "")
	for i, j := len(chars)-1, 0; i >= 0; i, j = i-1, j+1 {
		asciiToInt := int(chars[i][0]) - removal
		num = num + (int(math.Pow(26, float64(j))) * asciiToInt)
	}
	return num
}

func main() {
	fmt.Println(excelNum("CA"))
}
