package a10

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		array []int
	} {
		{
			name: "case1",
			array: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		},
		{
			name: "case2",
			array: []int{1, 1},
		},
		{
			name: "case3",
			array: []int{4, 3, 2, 1, 4},
		},
		{
			name: "case4",
			array: []int{1, 2, 1},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%v\n", tt.array)
			f(tt.array)
		})
	}
}
