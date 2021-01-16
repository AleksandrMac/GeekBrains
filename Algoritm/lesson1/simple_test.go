package lesson1

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type maxTest struct {
	Name        string
	arg1        []int32
	expected    int32
	expectedErr error
}

var maxTests = []maxTest{
	{"test1", []int32{}, -1, errors.New("please enter one or more numbers")},
	{"test2", []int32{4, 8}, 8, nil},
	{"test3", []int32{-10, 0, 4, 8}, 8, nil},
	{"test4", []int32{50, 4, -1, -10, 0, 4, 8}, 50, nil},
}

func TestMax(t *testing.T) {

	for _, test := range maxTests {
		output, outputErr := Max(test.arg1)
		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
		if outputErr != nil {
			if outputErr.Error() != test.expectedErr.Error() {
				t.Errorf("Output %q not equal to expected %q", outputErr, test.expectedErr)
			}
		}
	}
}

type lenNumTest struct {
	Name     string
	arg1     int64
	expected uint16
}

var lenNumTests = []lenNumTest{
	{"test1", 0, 1},
	{"test2", 1, 1},
	{"test3", -1, 1},
	{"test4", 10, 2},
	{"test5", -10, 2},
	{"test6", 9, 1},
	{"test7", 9, 1},
	{"test8", 11, 2},
	{"test9", 11, 2},
	{"test10", 111, 3},
	{"test11", 99, 2},
	{"test12", 100, 3},
	{"test13", 999, 3},
	{"test14", 1000, 4},
}

func TestLenNum(t *testing.T) {

	for _, test := range lenNumTests {
		output := lenNum(test.arg1)
		if output != test.expected {
			t.Errorf("%q: Output %d not equal to expected %d", test.Name, output, test.expected)
		}
	}
}

func TestAutomorphicNumbers(t *testing.T) {
	output := AutomorphicNumbers(10000)
	expected := []uint32{0, 1, 5, 6, 25, 76, 376, 625, 9376}
	fmt.Println(output)
	fmt.Println(expected)
	if reflect.DeepEqual(output, expected) {
		t.Errorf("Output %q not equal to expected %q", output, expected)
	}
}
