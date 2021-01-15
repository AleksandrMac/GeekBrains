package lesson1

import (
	"errors"
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
