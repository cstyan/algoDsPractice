package main

// started 15:47
// done 15:59

import (
	"fmt"
	"strings"
)

type position struct {
	x int
	y int
}

var initPosition = position{0, 0}

func moveRobot(pos *position, command string) {
	switch command {
	case "u":
		pos.y = pos.y + 1
	case "d":
		pos.y = pos.y - 1
	case "r":
		pos.x = pos.x + 1
	case "l":
		pos.x = pos.x - 1
	default:
		//unknown
	}
}

// return whether or robot does a circle
func circle(input string) bool {
	var robotPosition = position{0, 0}
	input = strings.ToLower(input)
	inputSplit := strings.Split(input, "")
	for _, val := range inputSplit {
		moveRobot(&robotPosition, val)
	}
	fmt.Println("final position: ", robotPosition)
	if robotPosition != initPosition {
		return false
	}
	return true
}

func main() {
	fmt.Println("circle? ", circle("udllurdrrd"))
}
