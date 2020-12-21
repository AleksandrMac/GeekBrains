package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	var a, b, res float32
	var op string
	fmt.Println("Введите первое число")

	validScan(&a)

	fmt.Println("Введите второе число")
	validScan(&b)

	fmt.Println("Введите арифметическую операцию:")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Деление на ноль не возможно.")
			break
		}
		res = a / b
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат: %f\n", res)
}

func validScan(param interface{}) (count int, err error) {
	for {
		count, err := fmt.Scanln(param)
		if err == nil && count == 1 {
			return count, err
		}
		return fmt.Println("Please input correct", reflect.TypeOf(param))
	}
}
