package task

import (
	"fmt"
	"sync"
)

type counter uint64

// Task2 -
// Реализуйте функцию для разблокировки мьютекса с помощью defer
func Task2() {
	var count counter
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(count *counter) {
			lock.Lock()
			defer lock.Unlock()
			count.add()
			wg.Done()
		}(&count)
	}
	wg.Wait()
	fmt.Println(count)
}

func (c *counter) add() {
	*c += counter(1)
}
