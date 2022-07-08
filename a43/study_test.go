package a43

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name   string
		matrix [][]int
		m      int
		n      int
	}{
		{
			name: "case1",
			m:    5,
			n:    4,
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
				{17, 18, 19, 20},
			},
		},
		{
			name: "case2",
			m:    1,
			n:    4,
			matrix: [][]int{
				{1, 2, 3, 4},
			},
		},
		{
			name: "case3",
			m:    4,
			n:    1,
			matrix: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},
		{
			name: "case4",
			m:    2,
			n:    4,
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
			},
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			f(item.matrix, item.m, item.n)
		})
	}
}
