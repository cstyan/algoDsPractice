package main

// start 15:10
// finished 15:42, had to google some golang things

import (
	"fmt"
)

type node struct {
	left  *node
	right *node
	value int
}

func (n *node) String() string {
	return fmt.Sprintf("node: %d left: %d right %d", n.value, n.left.value, n.right.value)
}

var input = [6]int{3, 2, 1, 6, 0, 5}

func findLargest(vals []int) (index int, largest int) {
	var curIndex, curVal = -1, -1
	for i, val := range vals {
		if val > curVal {
			curVal = val
			curIndex = i
		}
	}
	return curIndex, curVal
}

func buildMaxBinTree(vals []int) *node {
	// exit early
	if len(vals) == 0 {
		// treat -1 as a non-existant leaf
		return &node{nil, nil, -1}
	}
	var currentNode node
	index, largest := findLargest(vals)
	currentNode.value = largest
	currentNode.left = buildMaxBinTree(vals[:index])
	currentNode.right = buildMaxBinTree(vals[index+1 : len(vals)])
	return &currentNode
}

func main() {
	fmt.Println("root node is: ", buildMaxBinTree(input[:]).String())
}
