package main

import (
	"fmt"
	"strconv"
)

// You're now a baseball game point recorder.

// Given a list of strings, each string can be one of the 4 following types:

// Integer (one round's score): Directly represents the number of points you get in this round.
// "+" (one round's score): Represents that the points you get in this round are the sum of the last two valid round's points.
// "D" (one round's score): Represents that the points you get in this round are the doubled data of the last valid round's points.
// "C" (an operation, which isn't a round's score): Represents the last valid round's points you get were invalid and should be removed.

// Each round's operation is permanent and could have an impact on the round before and the round after.

// You need to return the sum of the points you could get in all the rounds.

func calcPoints(array []string) int {
	lastValidRound := 0
	secondLastValidRound := 0
	pointSum := 0
	for _, item := range array {
		switch item {
		case "+":
			addValue := lastValidRound + secondLastValidRound
			pointSum = pointSum + addValue
			secondLastValidRound = lastValidRound
			lastValidRound = addValue
		case "D":
			addValue := lastValidRound * 2
			pointSum = pointSum + (lastValidRound * 2)
			secondLastValidRound = lastValidRound
			lastValidRound = addValue
		case "C":
			pointSum = pointSum - lastValidRound
			lastValidRound = secondLastValidRound
			secondLastValidRound = 0
		default:
			intVal, err := strconv.Atoi(item)
			if err != nil {
				fmt.Println("Invalid character: ", err)
			}
			pointSum = pointSum + intVal
			secondLastValidRound = lastValidRound
			lastValidRound = intVal
		}
		fmt.Printf("item was: %s and sum is %d\n", item, pointSum)
	}
	return pointSum
}

func main() {
	arrayOne := [...]string{"5", "2", "C", "D", "+"}
	fmt.Println("sum one: ", calcPoints(arrayOne[:]))
	arrayTwo := [...]string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println("sum two: ", calcPoints(arrayTwo[:]))
}
