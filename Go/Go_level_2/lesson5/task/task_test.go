package task

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
)

type Procent struct {
	write, read int
}

func BenchmarkTask3Mutex(b *testing.B) {
	data := []Procent{
		{write: 10, read: 90},
		{write: 50, read: 50},
		{write: 90, read: 10},
	}

	for _, val := range data {
		name := fmt.Sprintf(" write:%v /read:%v", val.write, val.read)
		b.Run(strconv.Itoa(runtime.GOMAXPROCS(4))+name, func(b *testing.B) {
			b.SetParallelism(100)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					Task3Mutex(val.write, val.read)
				}
			})

		})
	}
}

func BenchmarkTask3RWMutex(b *testing.B) {
	data := []Procent{
		{write: 10, read: 90},
		{write: 50, read: 50},
		{write: 90, read: 10},
	}
	for _, val := range data {
		name := fmt.Sprintf(" write:%v /read:%v", val.write, val.read)
		b.Run(strconv.Itoa(runtime.GOMAXPROCS(4))+name, func(b *testing.B) {
			b.SetParallelism(100)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					Task3RWMutex(val.write, val.read)
				}
			})
		})
	}
}
