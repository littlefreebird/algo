package a6

import (
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct{
		name string
		s1 string
		s2 string
	} {
		{
			name: "case1",
			s1: "abcdef",
			s2: "abc",
		},
		{
			name: "case2",
			s1: "abcdef",
			s2: "cab",
		},
		{
			name: "case3",
			s1: "abcdef",
			s2: "fe",
		},
		{
			name: "case4",
			s1: "abcxyzdef",
			s2: "z",
		},
		{
			name: "case5",
			s1: "abcdef",
			s2: "g",
		},
		{
			name: "case6",
			s1: "",
			s2: "",
		},
		{
			name: "case7",
			s1: "",
			s2: "abc",
		},
		{
			name: "case8",
			s1: "abc",
			s2: "",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%s, %s => %s\n", tt.s1, tt.s2, f(tt.s1, tt.s2))
		})
	}
}
