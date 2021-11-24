package a7

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		s string
		n int
	}{
		{
			name: "case1",
			s: "abcdefghijki",
			n: 3,
		},
		{
			name: "case2",
			s: "a",
			n: 1,
		},
		{
			name: "case3",
			s: "ab",
			n: 1,
		},
		{
			name: "case4",
			s: "ab",
			n: 2,
		},
		{
			name: "case5",
			s: "abc",
			n: 1,
		},
		{
			name: "case6",
			s: "abcdefg",
			n: 1,
		},
		{
			name: "case7",
			s: "abc",
			n: 2,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%s\n", f(tt.s, tt.n))
		})
	}
}
