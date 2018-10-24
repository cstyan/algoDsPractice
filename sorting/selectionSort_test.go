package sort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	var cases = []struct {
		input     []int
		expected  []int
		increment bool
	}{
		{
			input:     []int{7, 5, 3, 10, 9, 45, 2, 34, 23},
			expected:  []int{2, 3, 5, 7, 9, 10, 23, 34, 45},
			increment: true,
		},
		{
			input:     []int{7, 5, 3, 10, 9, 45, 2, 34, 23},
			expected:  []int{45, 34, 23, 10, 9, 7, 5, 3, 2},
			increment: false,
		},
	}
	for _, c := range cases {
		sorted := c.input[:]
		selectionSort(c.increment, sorted)
		if !reflect.DeepEqual(c.expected, sorted) {
			t.Errorf("not sorted properly:\n\tinput: %+v\n\texpected: %v\n\toutput: %v\n", c.input, c.expected, sorted)
		}
	}
}
