package a23

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		arr []int
		size int
	} {
		{
			name: "case1",
			arr: []int{},
			size: 0,
		},
		{
			name: "case2",
			arr: []int{1},
			size: 1,
		},
		{
			name: "case3",
			arr: []int{1,1},
			size: 1,
		},
		{
			name: "case4",
			arr: []int{1,2},
			size: 2,
		},
		{
			name: "case5",
			arr: []int{1, 1, 2},
			size: 2,
		},
		{
			name: "case7",
			arr: []int{1, 1, 2, 2},
			size: 2,
		},
		{
			name: "case8",
			arr: []int{1, 2, 2, 2, 3, 3, 3, 4, 5, 6, 6, 6, 7, 9, 10, 10, 10},
			size: 9,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			ret := f(item.arr)
			if ret != item.size {
				t.Fail()
			}
		})
	}
}