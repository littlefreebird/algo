package a49

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		a    []int
		k    int
	}{
		{
			name: "case1",
			a:    []int{1},
			k:    2,
		},
		{
			name: "case2",
			a:    []int{1, 2, 3, 4, 5, 6},
			k:    15,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			f(item.a, item.k)
		})
	}
}
