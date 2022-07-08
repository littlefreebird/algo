package a42

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
	}{
		{
			name: "case1",
			arr:  []int{1},
		},
		{
			name: "case2",
			arr:  []int{1, 2},
		},
		{
			name: "case3",
			arr:  []int{1, -2, 10, -2, -1, 5, -4},
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			f(item.arr)
		})
	}
}
