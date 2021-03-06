package main

import (
	"fmt"

	"github.com/AleksandrMac/GeekBrains/Go/Go_level_2/lesson5/task"
)

//nolint:gomnd
func main() {
	fmt.Println("----------Task1-------------")
	task.Task1()

	fmt.Println("----------Task2-------------")
	task.Task2()

	fmt.Println("----------Task3-------------")
	task.Task3Mutex(50, 50)
	task.Task3RWMutex(50, 50)
}
