package lesson3

import "fmt"

//Александр Макалов
//1. Попробовать оптимизировать пузырьковую сортировку.
//Сравнить количество операций сравнения оптимизированной и не оптимизированной программы.
//Написать функции сортировки, которые возвращают количество операций.

//Main -
func Main() {
	slice := []int32{5, -2, 5, 3, 9, 8, 6, 2, 1, 0, 6, 2, -1}

	sortSlice := BubbleSort(slice)
	fmt.Println("unsort", slice)
	fmt.Println("bubble sort", sortSlice)

	sortSlice = BubbleSortOptimum(slice)
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
	count := int32(0)
	for i := 0; i < len(sortedSlice); i++ {
		for j := range sortedSlice[:len(sortedSlice)-1] {
			if sortedSlice[j] > sortedSlice[j+1] {
				sortedSlice[j], sortedSlice[j+1] = sortedSlice[j+1], sortedSlice[j]
			}
			count++
		}
	}
	fmt.Println("\nBubbleSort - count = ", count)
	return sortedSlice
}

//BubbleSortOptimum -
func BubbleSortOptimum(slice []int32) []int32 {
	sortedSlice := make([]int32, len(slice))
	copy(sortedSlice, slice)
	count := int32(0)
	for i := 0; i < len(sortedSlice); i++ {
		for j := range sortedSlice[i : len(sortedSlice)-1] {
			if sortedSlice[j] > sortedSlice[j+1] {
				sortedSlice[j], sortedSlice[j+1] = sortedSlice[j+1], sortedSlice[j]
			}
			count++
		}
	}
	fmt.Println("\nBubbleSortOptimum - count = ", count)
	return sortedSlice
}

//InsertSort -
func InsertSort(slice []int32) []int32 {
	sortedSlice := make([]int32, len(slice))
	count := int32(0)

	for i := 0; i < len(slice); i++ {
		key := slice[i]
		j := i
		for ; j > 0 && sortedSlice[j-1] > key; j-- {
			sortedSlice[j] = sortedSlice[j-1]
			count++
		}
		sortedSlice[j] = key
	}
	fmt.Println("\nInsertSort - count =", count)
	return sortedSlice
}
