package a60

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		k    int
		c    int
	}{
		{
			name: "case1",
			arr:  []int{1, 2, 3, 3, 3, 4, 5, 6},
			k:    3,
			c:    3,
		},
		{
			name: "case2",
			arr:  []int{1, 2, 3},
			k:    3,
			c:    1,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			if item.c != f(item.arr, item.k) {
				t.Fail()
			}
		})
	}
}
