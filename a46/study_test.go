package a46

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name      string
		intervals []section
	}{
		{
			name: "case1",
			intervals: []section{
				{start: 10, end: 15},
				{start: 1, end: 5},
				{start: 2, end: 5},
				{start: 7, end: 9},
				{start: 2, end: 4},
			},
		},
		{
			name: "case2",
			intervals: []section{
				{start: 2, end: 3},
				{start: 5, end: 6},
				{start: 4, end: 5},
			},
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			f(item.intervals)
		})
	}
}
