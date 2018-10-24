package sort

// standard solution using indicies
func merge(left, right []int) []int {
	leftIndex := 0
	rightIndex := 0
	lenLeft := len(left)
	lenRight := len(right)
	var merged []int
	for leftIndex < lenLeft && rightIndex < lenRight {
		if left[leftIndex] < right[rightIndex] {
			merged = append(merged, left[leftIndex])
			leftIndex = leftIndex + 1
		} else {
			merged = append(merged, right[rightIndex])
			rightIndex = rightIndex + 1
		}
	}
	if leftIndex < lenLeft {
		merged = append(merged, left[leftIndex:]...)
	}
	if rightIndex < lenRight {
		merged = append(merged, right[rightIndex:]...)
	}
	return merged
}

// splits current array in half, determines if halves are ready to be merged
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// len(arr)/2 rounds down
	// arr[:num] is exlusive, does not include element at index num
	// arr[num:] is inclusive, does include element at index num
	left := arr[:len(arr)/2]
	right := arr[len(arr)/2:]
	if len(left) != 1 {
		left = mergeSort(left)
	}
	if len(right) != 1 {
		right = mergeSort(right)
	}
	return merge(left, right)
}
