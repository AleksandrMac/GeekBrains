package main

import "fmt"

func main() {
	fiboMap := make(map[int64]int64)
	var n int64
	fmt.Scanln(&n)

	for i := int64(0); i <= int64(n); i++ {
		fiboMap[i] = fibo(int64(i))
	}
	fmt.Println(fiboMap)
}

func fibo(n int64) int64 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibo(n-1) + fibo(n-int64(2))
	}
}
