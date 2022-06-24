package a24

import "testing"

func TestF(t *testing.T)  {
	cases := []struct{
		name string
		arr []int
		k int
		size int
	} {
		{
			name: "case1",
			arr: []int{},
			k: 1,
			size:0,
		},
		{
			name: "case2",
			arr: []int{1},
			k: 2,
			size: 1,
		},
		{
			name: "case3",
			arr: []int{1,2},
			k: 1,
			size: 1,
		},
		{
			name: "case4",
			arr: []int{1,2},
			k: 2,
			size:1,
		},
		{
			name: "case5",
			arr: []int{1,1, 2, 1, 3,4, 5, 1, 1},
			k:1,
			size: 4,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			ret := f(item.arr, item.k)
			if ret != item.size {
				t.Fail()
			}
		})
	}
}