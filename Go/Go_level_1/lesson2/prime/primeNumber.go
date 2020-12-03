package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("please, input topLimit fot prime number")
	var topLimit int64
	fmt.Scanln(&topLimit)
	var i int64
	for i = 0; i <= topLimit; i++ {
		if isPrimeNumber(i) {
			fmt.Println(i)
		}
	}
}

func isPrimeNumber(num int64) bool {
	if num == 1 || num == 0 {
		return false
	}
	for i := 2; math.Pow(float64(i), 2) <= float64(num); i++ {
		if int64(num)%int64(i) == 0 {
			return false
		}
	}
	return true
}
