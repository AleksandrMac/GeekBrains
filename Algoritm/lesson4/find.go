package lesson4

import "fmt"

//Horse -
func Horse(m int8) [][]int32 {

	var slice [][]int32
	slice = make([][]int32, m, m)
	for i := int8(0); i < m; i++ {
		slice[i] = make([]int32, m)
	}

	for i := int8(0); i < m; i++ {
		fmt.Println(slice[i])
	}
	return slice
}
