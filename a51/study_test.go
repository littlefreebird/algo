package a51

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct {
		name string
		arr  [][]int
		k    int
		m    int
		n    int
	}{
		{
			name: "case1",
			arr: [][]int{
				{1, 2, 4, 6, 7},
				{3, 5, 7, 9, 10},
				{4, 5, 8, 12, 15},
				{6, 9, 10, 30, 100},
			},
			m: 4,
			n: 5,
			k: 12,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			fmt.Println(f(item.arr, 0, item.n-1, 0, item.m-1, item.k))
		})
	}
}
