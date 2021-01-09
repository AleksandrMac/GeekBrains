package sort

import (
	"fmt"
)

func main() {
	slice := []int32{5, -2, 5, 3, 9, 8, 6, 2, 1, 0, 6, 2, -1}

	sortSlice := BubbleSort(slice)
	fmt.Println("unsort", slice)
	fmt.Println("bubble sort", sortSlice)

	sortSlice = InsertSort(slice)
	fmt.Println("unsort", slice)
	fmt.Println("insert sort", sortSlice)

}

//BubbleSort -
func BubbleSort(slice []int32) []int32 {
	sortedSlice := make([]int32, len(slice))
	copy(sortedSlice, slice)
	for i := 0; i < len(sortedSlice); i++ {
		for i := range sortedSlice[:len(sortedSlice)-1] {
			if sortedSlice[i] > sortedSlice[i+1] {
				sortedSlice[i], sortedSlice[i+1] = sortedSlice[i+1], sortedSlice[i]
			}
		}
	}
	return sortedSlice

}

// InsertSort -
func InsertSort(slice []int32) []int32 {
	sortedSlice := make([]int32, len(slice))

	for i := 0; i < len(slice); i++ {
		key := slice[i]
		j := i
		for ; j > 0 && sortedSlice[j-1] > key; j-- {
			sortedSlice[j] = sortedSlice[j-1]
		}
		sortedSlice[j] = key
	}
	return sortedSlice
}
