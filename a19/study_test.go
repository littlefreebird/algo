package a19

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		s string
	} {
		{
			name: "case1",
			s: "()",
		},
		{
			name: "case2",
			s: "{}{}[]()()",
		},
		{
			name: "case3",
			s: "([)]",
		},
		{
			name: "case4",
			s: "[]{}",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%s -> %t\n", tt.s, f(tt.s))
		})
	}
}
