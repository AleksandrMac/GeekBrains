package main

import "fmt"

// fibo структура хранящая в себе числа фибоначи
type fibo struct {
	mapa map[int64]int64
}

func main() {
	fi := fibo{
		mapa: map[int64]int64{},
	}

	var n int64
	fmt.Scanln(&n)

	for i := int64(0); i <= n; i++ {
		fmt.Println(i, " - ", fi.Fibo(i))
	}
}

// Fibo получаем значение числа фибоначи
func (f fibo) Fibo(n int64) int64 {
	number, ok := f.mapa[n]

	if ok {
		return number
	}
	switch n {
	case 0:
		f.mapa[0] = 0
		return 0
	case 1:
		f.mapa[0] = 0
		f.mapa[1] = 1
		return 1
	default:
		// если в мапе нет требуемого фибо,
		// происходит до заполнение мапы до требуемого значения
		if f.mapa[n-1] > 0 {
			f.mapa[n] = f.mapa[n-1] + f.mapa[n-2]
			return f.mapa[n]
		}
		return f.Fibo(n-1) + f.Fibo(n-2)
	}
}
