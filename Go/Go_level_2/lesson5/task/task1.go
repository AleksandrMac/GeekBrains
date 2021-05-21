package task

import (
	"fmt"
	"sync"
)

// Task1 -
// Напишите программу, которая запускает n потоков и дожидается завершения их всех
func Task1() {
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
}
