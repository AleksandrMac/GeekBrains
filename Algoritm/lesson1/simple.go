package lesson1

//Александр Макалов
//
import "errors"

//12. Написать функцию нахождения максимального из трех чисел.

//Max - return Max number from slice
func Max(a []int32) (int32, error) {
	if len(a) < 1 {
		return -1, errors.New("please enter one or more numbers")
	}
	max := a[0]
	var i int32
	for _, i = range a[1:] {
		if i > max {
			max = i
		}
	}
	return max, nil
}
