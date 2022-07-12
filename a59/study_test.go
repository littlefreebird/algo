package a59

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
	}{
		{
			name: "case1",
			arr:  []int{2, 1, 5, 3, 4},
		},
		{
			name: "case2",
			arr:  []int{2, 3, 1},
		},
	}
	for _, item := range cases {
		f(item.arr)
	}
}
