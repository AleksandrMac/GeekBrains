package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, res float32
	var op string
	fmt.Println("Введите первое число")

	scanNumber(&a)

	fmt.Println("Введите второе число")
	scanNumber(&b)

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
		res = a / b
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат: %f\n", res)
}

func scanNumber(param interface{}) {
	for {
		count, err := fmt.Scanln(param)
		if err == nil && count == 1 {
			break
		} else {
			fmt.Println("Please input correct number")
		}
	}
}
