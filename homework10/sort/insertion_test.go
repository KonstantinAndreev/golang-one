package sort

import "testing"

var flagtests = []struct {
	in  []int
	out []int
}{
	{
		in:  []int{1, 3, 4, 2, 0},
		out: []int{0, 1, 2, 3, 4},
	},
	{
		in:  []int{1, 8, 9, 2, 0, 6},
		out: []int{0, 1, 2, 6, 8, 9},
	},
}

func TestInsertionSort(t *testing.T) {
	for _, tt := range flagtests {
		realOut := InsertionSort(tt.in)
		for i, val := range realOut {
			if val != tt.out[i] {
				t.Errorf("got %q, want %q", val, tt.out)
			}
		}
	}
}
