package a27

import (
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct {
		name string
		a    []int
		r    []int
	}{
		{
			name: "case1",
			a:    []int{1},
			r:    []int{1},
		},
		{
			name: "case2",
			a:    []int{1, 2, 3},
			r:    []int{1, 3, 2},
		},
		{
			name: "case3",
			a:    []int{5, 4, 3, 2, 1},
			r:    []int{1, 2, 3, 4, 5},
		},
		{
			name: "case4",
			a:    []int{1, 2, 5, 4, 3},
			r:    []int{1, 3, 2, 4, 5},
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			r := f(item.a)
			equal := true
			for idx, _ := range r {
				if r[idx] != item.r[idx] {
					equal = false
					break
				}
			}
			if !equal {
				t.Fail()
			}
		})
	}
}
