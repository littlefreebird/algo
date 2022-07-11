package a53

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
	}{
		{
			name: "case1",
			arr:  []int{1, 3, 2, 5, 4, 7, 9, 10},
		},
		{
			name: "case2",
			arr:  []int{1, 3, 5, 7},
		},
		{
			name: "case3",
			arr:  []int{2, 4, 6, 8},
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			f(item.arr)
		})
	}
}
