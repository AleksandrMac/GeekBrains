package main

import "fmt"

func main() {
	//first()
	second()
}

func first() {
	for i := 1; i <= 100; i++ {

		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func second() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
		} else {
			if i%5 == 0 {
				fmt.Println("Buzz")
			} else {
				fmt.Println(i)
			}
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println()
		}
	}
}
