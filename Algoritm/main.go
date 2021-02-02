package main

import (
	"fmt"

	"github.com/AleksandrMac/GeekBrains/algoritm/lesson4"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(lesson1.Rand100Default(), "\t", lesson1.Rand100Custom())
	// }
	//fmt.Println(lesson1.AutomorphicNumbers(10000))

	slice := lesson4.Horse(8)
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
