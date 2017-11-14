package main

import (
	"fmt"
	"math/rand"
	// "errors"
)

type pair struct {
	val1 int
	val2 int
}

// O(n)
func listToMap(list []int) map[int]struct{} {
	var empty struct{}
	m := make(map[int]struct{})
	// we only care about unique #'s so lets ignore duplicates for now
	// and just ensure we know about the existence of each unique # in the list
	for _, element := range list {
		m[element] = empty
	}
	return m
}

// find if there are elements in two lists that add up to k
// return all unique pairs of elements that do
func twoSum(listOne []int, listTwo []int, num int) {
	if len(listOne) == 0 || len(listTwo) == 0 {
		// return errors.New("no pairs, one of the input lists has no elements")
	}
	fmt.Printf("sum of numbers in two arrays that add up to: %d\n\n", num)
	fmt.Printf("first array: \"%v\"\n\n", listOne)
	fmt.Printf("second array: \"%v\"\n\n", listTwo)
	mapOne := listToMap(listOne)
	mapTwo := listToMap(listTwo)
	fmt.Println("pairs:")
	for key := range mapOne {
		if _, ok := mapTwo[num-key]; ok == true {
			fmt.Printf("\t%d, %d\n", key, num-key)
		}
	}
}

func main() {
	size := 100
	testRange := 100
	listOne := []int{}
	listTwo := []int{}

	for i := 0; i < size; i++ {
		listOne = append(listOne, rand.Intn(testRange))
	}
	for j := 0; j < size; j++ {
		listTwo = append(listTwo, rand.Intn(testRange))
	}
	twoSum(listOne[:], listTwo[:], 42)
}
