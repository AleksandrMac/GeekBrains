package lesson3

//Александр Макалов
//1. Попробовать оптимизировать пузырьковую сортировку.
//Сравнить количество операций сравнения оптимизированной и не оптимизированной программы.
//Написать функции сортировки, которые возвращают количество операций.

/*
func main() {
	slice := []int32{5, -2, 5, 3, 9, 8, 6, 2, 1, 0, 6, 2, -1}

	sortSlice := bubbleSort(slice)
	fmt.Println("unsort", slice)
	fmt.Println("bubble sort", sortSlice)

	sortSlice = insertSort(slice)
	fmt.Println("unsort", slice)
	fmt.Println("insert sort", sortSlice)

}*/

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

func insertSort(slice []int32) []int32 {
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
