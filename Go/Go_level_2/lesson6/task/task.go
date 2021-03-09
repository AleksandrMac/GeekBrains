package task

import (
	"os"
	"runtime"
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
		go func(count *uint32, lock *sync.Mutex) {
			lock.Lock()
			defer lock.Unlock()
			*count++
		}(&count, &lock)
	}
}

// Task2 - Написать многопоточную программу, в которой будет
// использоваться явный вызов планировщика. Выполните трассировку программы
func Task2() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	var count uint32
	lock := sync.Mutex{}

	for i := 0; i < num; i++ {
		go func(count *uint32, lock *sync.Mutex) {
			lock.Lock()
			defer lock.Unlock()
			*count++
		}(&count, &lock)
		if i%1e2 == 0 {
			runtime.Gosched()
		}
	}
}
