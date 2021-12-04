package a17

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		array []int
		target int
	} {
		{
			name: "case1",
			array: []int{1, 2, 3, 4, 5, 6, 7, 7, 8},
			target: 16,
		},
		{
			name: "case2",
			array: []int{1, 5, 6, 7},
			target: 19,
		},
		{
			name: "case3",
			array: []int{3, 4, 5 ,6, 7, 8, 9},
			target: 24,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			f(tt.array, tt.target)
		})
	}
}
