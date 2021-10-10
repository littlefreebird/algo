package main

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		p1 []int
		p2 int
	} {
		{
			name: "case1",
			p1 : []int{1, 3, 5, 9},
			p2 : 8,
		},
		{
			name : "case2",
			p1 : []int{},
			p2 : 8,
		},
		{
			name : "case3",
			p1 : []int{2, 3, 4, 7},
			p2 : 11,
		},
		{
			name : "case4",
			p1 : []int{5, 6, 7, 8},
			p2 : 8,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			f(tt.p1, tt.p2)
		})
	}
}