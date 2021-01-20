package lesson2

import "testing"

type calcPerformerTest struct {
	Name     string
	start    int8
	end      int8
	expected int64
}

var calcPerformerTests = []calcPerformerTest{
	{"test6", 2, 6, 3},
	{"test1", 2, 7, 3},
	{"test2", 2, 2, 0},
	{"test3", 2, 3, 1},
	{"test4", 2, 4, 2},
	{"test5", 2, 5, 2},
	{"test7", 3, 20, 32},
}

func TestCalcPerformer(t *testing.T) {
	for _, test := range calcPerformerTests {
		output := CalcPerformer(test.start, test.end, true)
		if output != test.expected {
			t.Errorf("%q: Output %d not equal to expected %d", test.Name, output, test.expected)
		}
	}
}
