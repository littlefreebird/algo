package a22

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		arr []int
	} {
		{
			name: "case1",
			arr: []int{},
		},
		{
			name: "case2",
			arr: []int{1},
		},
		{
			name: "case3",
			arr: []int{1, 2},
		},
		{
			name: "case4",
			arr: []int{1, 2, 3},
		},
		{
			name: "case5",
			arr: []int{1, 2, 3, 4},
		},
		{
			name: "case6",
			arr: []int{1, 2, 3, 4, 5},
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			f(item.arr)
		})
	}
}
