package main

import (
	"fmt"

	"github.com/AleksandrMac/GeekBrains/algoritm/lesson1"
	"github.com/AleksandrMac/GeekBrains/algoritm/lesson3"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(lesson1.Rand100Default(), "\t", lesson1.Rand100Custom())
	// }
	fmt.Println(lesson1.AutomorphicNumbers(10000))

	lesson3.Main()
}
