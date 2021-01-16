package lesson1

//Александр Макалов
//
import (
	"errors"
	"math"
	"math/rand"
	"time"
)

//12. Написать функцию нахождения максимального из трех чисел.

//Max - return Max number from slice
func Max(a []int32) (int32, error) {
	if len(a) < 1 {
		return -1, errors.New("please enter one or more numbers")
	}
	max := a[0]
	for _, i := range a[1:] {
		if i > max {
			max = i
		}
	}
	return max, nil
}

//13. * Написать функцию, генерирующую случайное число от 1 до 100.
//	- с использованием стандартной функции rand()

//Rand100Default - return random number (1, 100)
func Rand100Default() uint64 {
	return (rand.Uint64())%100 + 1
}

//	- без использования стандартной функции rand()

//Rand100Custom - return random number (1, 100)
func Rand100Custom() uint64 {
	startN := time.Now().Nanosecond()
	return uint64(3*startN/1234+1)%100 + 1
}

//14. * Автоморфные числа. Натуральное число называется автоморфным, если оно равно последним
// цифрам своего квадрата. Например, 25 \ :sup: `2` = 625. Напишите программу, которая вводит
// натуральное число N и выводит на экран все автоморфные числа, не превосходящие N.

//AutomorphicNumbers -
func AutomorphicNumbers(top uint32) []uint32 {
	slice := []uint32{}
	for i := uint32(0); i < top; i++ {
		pow2 := i * i
		l := lenNum(int64(i))
		if i == pow2%uint32(math.Pow(10, float64(l))) {
			slice = append(slice, i)
		}
	}
	return slice
}

//LenNum -
func lenNum(num int64) (len uint16) {
	num = int64(math.Abs(float64(num)))
	len = 1
	for num >= 10 {
		num = num / 10
		len++
	}
	return len
}
