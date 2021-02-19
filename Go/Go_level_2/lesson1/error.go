package main

import "fmt"

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()

	var a int
	a = 1 / a
}
