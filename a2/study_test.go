package a2

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		p1 int
		p2 int
	} {
		{
			name: "case1",
			p1: 123,
			p2: 456,
		},
		{
			name: "case2",
			p1: 35,
			p2: 45,
		},
		{
			name: "case3",
			p1: 467,
			p2: 3,
		},
		{
			name: "case4",
			p1: 0,
			p2: 123,
		},
		{
			name: "case5",
			p1: 0,
			p2: 0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			f(tt.p1, tt.p2)
		})
	}
}
