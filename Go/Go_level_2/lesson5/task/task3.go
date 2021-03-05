package task

import "sync"

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
		go func(count *counter) {
			lock.Lock()
			defer lock.Unlock()
			count.add()
		}(&count)
	}

	for i := 0; i < num*r; i++ {
		go func(count *counter) {
			lock.Lock()
			defer lock.Unlock()
			_ = *count
		}(&count)
	}
}

// Task3RWMutex -
func Task3RWMutex(w, r int) {
	var count counter
	lock := sync.RWMutex{}

	for i := 0; i < num*w; i++ {
		go func(count *counter) {
			lock.Lock()
			defer lock.Unlock()
			count.add()
		}(&count)
	}

	for i := 0; i < num*r; i++ {
		go func(count *counter) {
			lock.RLock()
			defer lock.RUnlock()
			_ = *count
		}(&count)
	}
}
