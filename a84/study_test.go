package a84

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		k    int
	}{
		{
			name: "case1",
			arr:  []int{25, 7, 4, 9, 10, 43, 5},
			k:    3,
		},
		{
			name: "case2",
			arr:  []int{1},
			k:    1,
		},
		{
			name: "case3",
			arr:  []int{3, 4, 5, 6, 7, 8},
			k:    2,
		},
		{
			name: "case4",
			arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			k:    5,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			f(item.arr, item.k)
		})
	}
}
