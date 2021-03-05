package main

import (
	"runtime"
	"strconv"
	"testing"

	"github.com/AleksandrMac/GeekBrains/Go/Go_level_2/lesson5/task"
)

type Procent struct {
	write, read int
}

func BenchmarkTask3Mutex(b *testing.B) {
	data := []Procent{
		{10, 90},
		{50, 50},
		{90, 10},
	}

	for _, val := range data {
		b.Run(strconv.Itoa(runtime.GOMAXPROCS(4)), func(b *testing.B) {
			b.SetParallelism(100)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					task.Task3Mutex(val.write, val.read)
				}
			})

		})
	}
}

func BenchmarkTask3RWMutex(b *testing.B) {
	data := []Procent{
		{10, 90},
		{50, 50},
		{90, 10},
	}
	for _, val := range data {
		b.Run(strconv.Itoa(runtime.GOMAXPROCS(4)), func(b *testing.B) {
			b.SetParallelism(100)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					task.Task3RWMutex(val.write, val.read)
				}
			})
		})
	}
}
