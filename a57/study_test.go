package a57

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
	}{
		{
			name: "case1",
			arr:  []int{0, 1, 3, 4, 1},
		},
		{
			name: "case2",
			arr:  []int{0, 3, 2, 1, 5, 5, 6, 7, 7, 9},
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			f(item.arr)
		})
	}
}
