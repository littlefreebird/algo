package a32

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		n    int
		r    string
	}{
		{
			name: "case1",
			n:    1,
			r:    "1",
		},
		{
			name: "case2",
			n:    2,
			r:    "11",
		},
		{
			name: "case3",
			n:    3,
			r:    "21",
		},
		{
			name: "case4",
			n:    4,
			r:    "1211",
		},
		{
			name: "case5",
			n:    5,
			r:    "111221",
		},
		{
			name: "case6",
			n:    6,
			r:    "312211",
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			if item.r != f(item.n) {
				t.Fail()
			}
		})
	}
}
