package a52

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		min  int
	}{
		{
			name: "case1",
			arr:  []int{4, 5, 6, 7, 1, 2, 3},
			min:  1,
		},
		{
			name: "case2",
			arr:  []int{3},
			min:  3,
		},
		{
			name: "case3",
			arr:  []int{3, 3},
			min:  3,
		},
		{
			name: "case4",
			arr:  []int{4, 4, 5, 5, 6, 6, 1, 1, 2, 2, 3},
			min:  1,
		},
		{
			name: "case5",
			arr:  []int{4, 4, 4, 1, 2, 3},
			min:  1,
		},
		{
			name: "case6",
			arr:  []int{4, 5, 6, 1, 1, 1},
			min:  1,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			if item.min != f(item.arr) {
				t.Fail()
			}
		})
	}
}
