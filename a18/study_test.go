package a18

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		array []int
		n int
	} {
		{
			name: "case1",
			array: []int{1, 2, 3, 4, 5, 6, 7},
			n: 2,
		},
		{
			name: "case2",
			array: []int{1, 2, 3},
			n: 1,
		},
		{
			name: "case3",
			array: []int{1, 2, 3},
			n: 4,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			f(tt.array, tt.n)
		})
	}
}
