package prime

import (
	"math"
)

/*func main() {
	fmt.Println("please, input topLimit fot prime number")
	var topLimit int64
	fmt.Scanln(&topLimit)
	for i := int64(0); i <= topLimit; i++ {
		if isPrimeNumber(i) {
			fmt.Println(i)
		}
	}
}*/

//IsPrimeNumber - возвращает истину когда num является простым числом, иначе ложь
func IsPrimeNumber(num int64) bool {
	if num < 2 {
		return false
	}
	// ищем числа квадрат которых не превосходит делимого, при делении на которые остаток будет равен нулю
	// если хотя бы одно такое число есть, то число не является простым
	for i := int64(2); math.Pow(float64(i), 2) <= float64(num); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
