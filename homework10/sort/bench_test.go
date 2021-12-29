package sort

import (
	"testing"
)

var x []int

func BenchmarkSort(b *testing.B) {
	var res []int
	testSlice := []int{1, 3, 4, 2, 0}
	for i := 0; i < b.N; i++ {
		res = InsertionSort(testSlice)
	}
	x = res
}
