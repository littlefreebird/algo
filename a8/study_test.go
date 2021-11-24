package a8

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		s string
	} {
		{
			name: "case1",
			s: "",
		},
		{
			name: "case2",
			s: "   abc",
		},
		{
			name: "case3",
			s: "   -123",
		},
		{
			name: "case4",
			s: "   4",
		},
		{
			name: "case5",
			s: "  -34t",
		},
		{
			name: "case6",
			s: "1234567abc",
		},
		{
			name: "case7",
			s: "-345abc",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%s -> %d\n", tt.s, f(tt.s))
		})
	}
}
