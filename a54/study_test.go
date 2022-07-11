package a54

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		r    int
	}{
		{
			name: "case1",
			arr:  []int{1, 2},
			r:    -1,
		},
		{
			name: "case2",
			arr:  []int{1, 2, 1},
			r:    1,
		},
		{
			name: "case3",
			arr:  []int{1, 1, 1, 1, 2, 2, 2},
			r:    1,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			if item.r != f(item.arr) {
				t.Fail()
			}
		})
	}
}
