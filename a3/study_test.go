package a3

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		s string
	} {
		{
			name: "case1",
			s: "abc",
		},
		{
			name: "case2",
			s: "abcdefea",
		},
		{
			name: "case3",
			s: "a",
		},
		{
			name: "case4",
			s: "abcbcdghg",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			f(tt.s)
		})
	}
}