package main

import "fmt"

// takes in an array of integers, finds the indicies of two integers
// whose sum is euqal to target
// problem states that we can assume there is only one solution per input
func twoSum(nums []int, target int) (int, int, error) {
	for index, element := range nums {
		// indexOne := index
		for internalIndex, internalElement := range nums {
			if internalIndex == index {
				continue
			}
			if (element + internalElement) == target {
				return index, internalIndex, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("No sum of two elements for: %d", target)

}

func main() {
	target := 17
	numbers := [6]int{2, 3, 7, 5, 11, 13}
	fmt.Println("Numbers available are: ", numbers)
	indexOne, indexTwo, err := twoSum(numbers[:], target)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fmt.Sprintf("The indicies for %d are: %d %d", target, indexOne, indexTwo))
}
