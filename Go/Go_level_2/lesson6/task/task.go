package task

import (
	"os"
	"runtime/trace"
	"sync"
)

const num = 1000

// Task1 - Написать программу, которая использует мьютекс для
// безопасного доступа к данным из нескольких потоков.
// Выполните трассировку программы
func Task1() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	var count uint32
	lock := sync.Mutex{}

	for i := 0; i < num; i++ {
		go func(count *go uint32, lock *sync.Mutex) {
			lock.Lock()
			defer lock.Unlock()
			*count++
		}(&count, &lock)
	}
}
