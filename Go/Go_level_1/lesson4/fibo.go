package main

import "fmt"

// структура хранящая в себе числа фибоначи
type fibo struct {
	mapa map[int64]int64
}

func main() {
	fi := fibo{
		mapa: map[int64]int64{},
	}

	var n int64
	fmt.Scanln(&n)

	for i := int64(0); i <= int64(n); i++ {
		fmt.Println(i, " - ", fi.Get(i))
	}
}

//получаем значение числа фибоначи
func (f fibo) Get(n int64) int64 {
	number, ok := f.mapa[n]

	if ok {
		return number
	}
	if n == 0 {
		f.mapa[0] = 0
		return 0
	}
	if n == 1 {
		f.mapa[1] = 1
		return 1
	}
	if f.mapa[n-1] > 0 {
		return f.mapa[n-1] + f.mapa[n-2]
	}
	// если в мапе нет требуемого фибо,
	// происходит до заполнение мапы до требуемого значения
	f.Set(n)
	//после заполнения получаем требуемое число
	return f.Get(n)
}
func (f fibo) Set(n int64) {
	switch n {
	case 0:
		f.mapa[0] = 0
	case 1:
		f.mapa[1] = 1
	default:
		if f.mapa[n-1] > 0 {
			f.mapa[n] = f.mapa[n-1] + f.mapa[n-2]
		} else {
			f.Set(n - 1)
		}
	}
}
