package a58

import (
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct {
		name string
		A    []int
	}{
		{
			name: "case1",
			A:    []int{1, 2},
		},
		{
			name: "case2",
			A:    []int{1, 2, 3},
		},
		{
			name: "case3",
			A:    []int{1, 2, 3, 4},
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			f(item.A)
		})
	}
}
