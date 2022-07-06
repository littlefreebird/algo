package a35

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		min  int
	}{
		{
			name: "case1",
			arr:  []int{1},
			min:  0,
		},
		{
			name: "case2",
			arr:  []int{1, 2},
			min:  1,
		},
		{
			name: "case3",
			arr:  []int{1, 0, 2},
			min:  -1,
		},
		{
			name: "case4",
			arr:  []int{2, 0, 1, 3, 2, 1, 5},
			min:  3,
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
