package sort

func swapElements(arr []int, indexOne int, indexTwo int) {
	temp := arr[indexOne]
	arr[indexOne] = arr[indexTwo]
	arr[indexTwo] = temp
	// fmt.Println("Array is now: ", arr)
}

// selection sort in incrementing or decrementing order
func selectionSort(increment bool, arr []int) {
	// loop through all indicies in the array
	for i := 0; i < len(arr); i++ {
		swapIndex := i
		// loop through all elements after i, assuming elements up to i are already sorted
		for j := i + 1; j < len(arr); j++ {
			if increment == false && arr[j] > arr[swapIndex] {
				swapIndex = j
			} else if increment == true && arr[j] < arr[swapIndex] {
				swapIndex = j
			}
		}
		swapElements(arr, i, swapIndex)
	}
}
