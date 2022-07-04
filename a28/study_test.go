package a28

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name   string
		arr    []int
		target int
		pos    int
	}{
		{
			name:   "case1",
			arr:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			target: 8,
			pos:    7,
		},
		{
			name:   "case2",
			arr:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			target: 1,
			pos:    0,
		},
		{
			name:   "case3",
			arr:    []int{6, 8, 9, 1, 3, 4, 5},
			target: 1,
			pos:    3,
		},
		{
			name:   "case4",
			arr:    []int{6, 8, 9, 1, 3, 4, 5},
			target: 9,
			pos:    2,
		},
		{
			name:   "case5",
			arr:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			target: 9,
			pos:    -1,
		},
		{
			name:   "case6",
			arr:    []int{6, 8, 9, 1, 3, 4, 5},
			target: 7,
			pos:    -1,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			if item.pos != f(item.arr, item.target) {
				t.Fail()
			}
		})
	}
}
