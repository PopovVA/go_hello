package sort

import (
	"reflect"
	"testing"
)

type sliceSortTestCase struct {
	input  []int64
	output []int64
}

func TestInsertionSort(t *testing.T) {

	testCases := []sliceSortTestCase{
		{
			[]int64{0, 50, 20, 40, 40, 0},
			[]int64{0, 0, 20, 40, 40, 50},
		},
		{
			[]int64{0, 20, 30, 40, 50, 60},
			[]int64{0, 20, 30, 40, 50, 60},
		},
	}

	for _, cse := range testCases {
		InsertionSort(cse.input)
		if !reflect.DeepEqual(cse.input, cse.output) {
			t.Fatalf("Result incorrect, expected: %d, actual: %d", cse.output, cse.input)
		}
	}
}
