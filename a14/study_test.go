package a14

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		array []int
	}{
		{
			name: "case1",
			array: []int{-1, 1, 0},
		},
		{
			name: "case2",
			array: []int{-1, 5, 4, 6},
		},
		{
			name: "case3",
			array: []int{-5, 4, 3, 2, -8, 4},
		},
		{
			name: "case4",
			array: []int{-5, 4, 2, 3, 2, 3},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%v\n", tt.array)
			f(tt.array)
		})
	}
}
