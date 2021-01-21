package lesson3

import (
	"reflect"
	"testing"
)

type sliceTest struct {
	Name        string
	inputSlice  []int32
	outputSlice []int32
}

var calcPerformerTests = []sliceTest{
	{"test1", []int32{5, -5, 0, -10, 10, 1, 2, 3}, []int32{-10, -5, 0, 1, 2, 3, 5, 10}},
	{"test2", []int32{5, -2, 5, 3, 9, 8, 6, 2, 1, 0, 6, 2, -1}, []int32{-2, -1, 0, 1, 2, 2, 3, 5, 5, 6, 6, 8, 9}},
	/*{"test3", []int32{}, []int32{}},
	{"test4", []int32{}, []int32{}},
	{"test5", []int32{}, []int32{}},
	{"test6", []int32{}, []int32{}},
	{"test7", []int32{}, []int32{}},
	{"test8", []int32{}, []int32{}},*/
}

func TestBubbleSort(t *testing.T) {
	for _, test := range calcPerformerTests {
		output := BubbleSort(test.inputSlice)
		if !reflect.DeepEqual(output, test.outputSlice) {
			t.Errorf("%q: Output %d not equal to expected %d", test.Name, output, test.outputSlice)
		}
	}
}
