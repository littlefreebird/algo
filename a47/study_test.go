package a47

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name      string
		intervals []section
		sec       section
	}{
		{
			name: "case1",
			intervals: []section{
				{start: 3, end: 4},
				{start: 6, end: 7},
				{start: 15, end: 20},
			},
			sec: section{start: 1, end: 2},
		},
		{
			name: "case2",
			intervals: []section{
				{start: 3, end: 4},
				{start: 6, end: 7},
				{start: 15, end: 20},
			},
			sec: section{start: 1, end: 3},
		},
		{
			name: "case3",
			intervals: []section{
				{start: 3, end: 4},
				{start: 6, end: 7},
				{start: 15, end: 20},
			},
			sec: section{start: 1, end: 21},
		},
		{
			name: "case4",
			intervals: []section{
				{start: 3, end: 4},
				{start: 6, end: 7},
				{start: 15, end: 20},
			},
			sec: section{start: 4, end: 5},
		},
		{
			name: "case5",
			intervals: []section{
				{start: 3, end: 4},
				{start: 6, end: 7},
				{start: 15, end: 20},
			},
			sec: section{start: 9, end: 10},
		},
		{
			name: "case6",
			intervals: []section{
				{start: 3, end: 4},
				{start: 6, end: 7},
				{start: 15, end: 20},
			},
			sec: section{start: 4, end: 5},
		},
		{
			name: "case7",
			intervals: []section{
				{start: 3, end: 4},
				{start: 6, end: 7},
				{start: 15, end: 20},
			},
			sec: section{start: 30, end: 100},
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			f(item.intervals, item.sec)
		})
	}
}
