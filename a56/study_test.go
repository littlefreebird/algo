package a56

import (
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
	}{
		{
			name: "case1",
			arr:  []int{2, 2, 4, 4, 6, 6, 3, 7, 9, 9},
		},
		{
			name: "case2",
			arr:  []int{1, 3},
		},
		{
			name: "case3",
			arr:  []int{5, 6, 6, 5, 8, 7, 7, 8, 23, 45},
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			f(item.arr)
		})
	}
}
