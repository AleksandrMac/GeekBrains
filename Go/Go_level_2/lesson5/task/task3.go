package task

import (
	"sync"
)

const num = 1

// Task3Mutex -
// Протестируйте производительность множества действительных чисел,
// безопасность которого обеспечивается sync.Mutex и sync.RWMutex
// для разных вариантов использования:
// 10% запись, 90% чтение;
// 50% запись, 50% чтение;
// 90% запись, 10% чтение
func Task3Mutex(w, r int) {
	var count counter
	lock := sync.Mutex{}

	for i := 0; i < num*w; i++ {
		go func(count *counter, lock *sync.Mutex) {
			lock.Lock()
			defer lock.Unlock()
			count.add()
		}(&count, &lock)
	}

	for i := 0; i < num*r; i++ {
		go func(count *counter, lock *sync.Mutex) {
			lock.Lock()
			defer lock.Unlock()
			_ = *count
		}(&count, &lock)
	}
}

// Task3RWMutex -
func Task3RWMutex(w, r int) {
	var count counter
	lock := sync.RWMutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < num*w; i++ {
		wg.Add(1)
		go func(count *counter, lock *sync.RWMutex) {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			count.add()
		}(&count, &lock)
	}

	for i := 0; i < num*r; i++ {
		wg.Add(1)
		go func(count *counter, lock *sync.RWMutex) {
			defer wg.Done()
			lock.RLock()
			defer lock.RUnlock()
			_ = *count
		}(&count, &lock)
	}
}
