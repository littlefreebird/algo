package a44

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		b    bool
	}{
		{
			name: "case1",
			arr:  []int{0},
			b:    true,
		},
		{
			name: "case2",
			arr:  []int{2, 0, 1, 3, 0, 0},
			b:    true,
		},
		{
			name: "case3",
			arr:  []int{2, 1, 0, 3},
			b:    false,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			if item.b != f(item.arr) {
				t.Fail()
			}
		})
	}
}
