package main

import (
	"fmt"
)

func main() {
	slice := []int32{5, 5, 3, 9, 8, 6, 2, 1, 0, 6, 2}
	sortSlice := bubbleSort(slice)
	fmt.Println("unsort", slice)
	fmt.Println("bubble sort", sortSlice)
	sortSlice = insertSort(slice)
	fmt.Println("unsort", slice)
	fmt.Println("insert sort", sortSlice)

}

func bubbleSort(slice []int32) []int32 {
	sort := make([]int32, len(slice))
	copy(sort, slice)
	for i := 0; i < len(sort); i++ {
		for i := range sort[:cap(sort)-1] {
			if sort[i] > sort[i+1] {
				sort[i], sort[i+1] = sort[i+1], sort[i]
			}
		}
	}
	return sort

}

func insertSort(slice []int32) []int32 {
	sort := make([]int32, 1, len(slice))

	for i := 0; i < len(slice); i++ {
		sort = append(sort, slice[i])
		for j := i; sort[j] > slice[i]; j-- {
			sort[j], sort[j+1] = sort[j+1], sort[j]
		}
	}
	return sort[1:]
}
