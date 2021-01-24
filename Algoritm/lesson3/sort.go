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

	sortSlice = ShakerSort(slice)
	fmt.Println("unsort", slice)
	fmt.Println("insert sort", sortSlice)

	sortSlice = []int32{}
	for i := int32(1); i <= 100; i++ {
		sortSlice = append(sortSlice, i)
	}
	fmt.Println("slice", sortSlice)
	fmt.Println("find position", FindBin(0, sortSlice))
	fmt.Println("find position", FindBin(1, sortSlice))
	fmt.Println("find position", FindBin(2, sortSlice))
	fmt.Println("find position", FindBin(3, sortSlice))
	fmt.Println("find position", FindBin(4, sortSlice))
	fmt.Println("find position", FindBin(5, sortSlice))
	fmt.Println("find position", FindBin(6, sortSlice))
	fmt.Println("find position", FindBin(7, sortSlice))
	fmt.Println("find position", FindBin(8, sortSlice))
	fmt.Println("find position", FindBin(9, sortSlice))
	fmt.Println("find position", FindBin(10, sortSlice))
	fmt.Println("find position", FindBin(11, sortSlice))
	fmt.Println("find position", FindBin(12, sortSlice))

	fmt.Println("find position", FindBin(49, sortSlice))
	fmt.Println("find position", FindBin(50, sortSlice))
	fmt.Println("find position", FindBin(51, sortSlice))
	fmt.Println("find position", FindBin(97, sortSlice))
	fmt.Println("find position", FindBin(98, sortSlice))
	fmt.Println("find position", FindBin(99, sortSlice))
	fmt.Println("find position", FindBin(100, sortSlice))
	fmt.Println("find position", FindBin(101, sortSlice))
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

//2. *Реализовать шейкерную сортировку

//ShakerSort -
func ShakerSort(slice []int32) []int32 {
	//можно оптимизировать в 2 раза, если за один проход будем искать и минимум и максимум,
	//и расставлять их соответственно в начало и конец. A(n/2 * n/2)
	sortedSlice := make([]int32, len(slice))
	copy(sortedSlice, slice)
	count := int32(0)

	for k := 0; k < len(sortedSlice)-1; k++ {
		i := k
		var mPostion int
		var j int
		//ищем максимальный эллемент, запоминаем его, меняем его с крайним справа эллементом
		for j = mPostion + 1; j < len(sortedSlice)-i; j++ {
			if sortedSlice[mPostion] < sortedSlice[j] {
				mPostion = j
			}
			count++
		}
		j--
		sortedSlice[mPostion], sortedSlice[j] = sortedSlice[j], sortedSlice[mPostion]
		i++
		mPostion = len(sortedSlice) - i - 1
		for j = mPostion - 1; j >= i-1; j-- {
			if sortedSlice[mPostion] > sortedSlice[j] {
				mPostion = j
			}
			count++
		}
		j++
		sortedSlice[mPostion], sortedSlice[j] = sortedSlice[j], sortedSlice[mPostion]
	}

	fmt.Println("\nShakerSort - count =", count)
	return sortedSlice
}

//3. Реализовать бинарный алгоритм поиска в виде функции, которой передается отсортированный массив.
//Функция возвращает индекс найденного элемента или -1, если элемент не найден.

//FindBin -
func FindBin(value int32, slice []int32) int32 {
	l := int32(len(slice))
	left := l / 2
	right := l - left
	mid := left
	for i := int32(1); mid >= 0 || mid <= int32(len(slice)-1); i++ {
		if value == slice[mid] {
			return mid
		}
		if mid == 0 || mid == int32(len(slice)-1) {
			return -1
		}

		if value < slice[mid] {
			l = left
			left = l / 2
			right = l - left
			mid = mid - right
			continue
		}
		l = right
		left = l / 2
		right = l - left
		mid = mid + left
	}

	return -1
}
