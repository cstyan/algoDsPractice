package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	input := []int{7, 5, 3, 10, 9, 45, 2, 34, 23}
	expected := []int{2, 3, 5, 7, 9, 10, 23, 34, 45}
	sorted := mergeSort(input)
	if !reflect.DeepEqual(expected, sorted) {
		t.Errorf("not sorted properly:\n\tinput: %+v\n\texpected: %v\n\toutput: %v\n", input, expected, sorted)
	}
}
