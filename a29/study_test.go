package a29

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
			arr:    []int{1, 2, 4, 6, 8, 9},
			target: 9,
			pos:    5,
		},
		{
			name:   "case2",
			arr:    []int{9, 7, 6, 5, 3, 2},
			target: 9,
			pos:    0,
		},
		{
			name:   "case3",
			arr:    []int{100, 59, 45, 32, 9, 8, 7},
			target: 50,
			pos:    2,
		},
		{
			name:   "case4",
			arr:    []int{8, 18, 45, 78, 100, 200, 500, 1000},
			target: 7,
			pos:    0,
		},
		{
			name:   "case5",
			arr:    []int{3, 5, 8, 9, 34, 56},
			target: 100,
			pos:    6,
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
