package gorutine

import (
	"fmt"
)

// Task1 - С помощью пула воркеров написать программу, которая запускает 1000 горутин,
// каждая из которых увеличивает число на 1. Дождаться завершения всех горутин и убедиться,
// что при каждом запуске программы итоговое число равно 1000.
func Task1() {
	ch := make(chan uint16)
	for range [1000]bool{} {
		go func() {
			val := <-ch
			fmt.Println(val)
			ch <- add(val)
		}()
	}
	ch <- 0
	fmt.Println("outside ", <-ch)
}

func add(x uint16) uint16 { return x + 1 }
