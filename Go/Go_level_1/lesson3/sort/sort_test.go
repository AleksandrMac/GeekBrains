package sort

import "testing"

func BenchmarkBubbleSort(b *testing.B) {
	//slice := []int32{}

	for i := 0; i < b.N; i++ {
		BubbleSort([]int32{1, 34, 21, 1235, 12, 51, 23, 5, 341, 235, 461235, 3436, 123})
	}

	//GlobalF = slice
}
func BenchmarkInsertSort(b *testing.B) {
	//slice := []int32{}

	for i := 0; i < b.N; i++ {
		InsertSort([]int32{1, 34, 21, 1235, 12, 51, 23, 5, 341, 235, 461235, 3436, 123})
	}

	//GlobalF = slice
}
