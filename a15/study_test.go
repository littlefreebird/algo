package a15

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		array []int
		target int
	} {
		{
			name: "case1",
			array: []int{1, 2, 3},
			target: 100,
		},
		{
			name: "case2",
			array: []int{-3, 9, 8, 7, 5},
			target: 11,
		},
		{
			name: "case3",
			array: []int{2, 3, 4, 5, 2, 8},
			target: 10,
		},
		{
			name: "case4",
			array: []int{4, 5, 6, 7, 2, 1},
			target: 19,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			f(tt.array, tt.target)
		})
	}
}
