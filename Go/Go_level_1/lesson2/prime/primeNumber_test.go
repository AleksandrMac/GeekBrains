package prime

import (
	"fmt"
	"testing"
)

func TestIsPrimeNumber(t *testing.T) {
	tests := []struct {
		input int64
		want  bool
	}{
		{input: -5, want: false},
		{input: -4, want: false},
		{input: -3, want: false},
		{input: 0, want: false},
		{input: 1, want: false},
		{input: 2, want: true},
		{input: 3, want: true},
		{input: 4, want: false},
		{input: 5, want: true},
		{input: 6, want: false},
		{input: 79, want: true},
		{input: 89, want: true},
		{input: 193, want: true},
		{input: 199, want: true},
		{input: 1000, want: false},
		{input: 2000, want: false},
	}
	for _, tc := range tests {
		if got := IsPrimeNumber(tc.input); got != tc.want {
			t.Fatalf("%v: expected: %v, got: %v", tc.input, tc.want, got)
		}
	}

}

func ExampleIsPrimeNumber() {
	number := int64(-5)
	fmt.Printf("Число %v является простым?, %v", number, IsPrimeNumber(number))

	number = int64(-1)
	fmt.Printf("Число %v является простым?, %v", number, IsPrimeNumber(number))

	number = int64(0)
	fmt.Printf("Число %v является простым?, %v", number, IsPrimeNumber(number))

	number = int64(1)
	fmt.Printf("Число %v является простым?, %v", number, IsPrimeNumber(number))

	number = int64(2)
	fmt.Printf("Число %v является простым?, %v", number, IsPrimeNumber(number))

	number = int64(4)
	fmt.Printf("Число %v является простым?, %v", number, IsPrimeNumber(number))

	number = int64(7)
	fmt.Printf("Число %v является простым?, %v", number, IsPrimeNumber(number))
}
