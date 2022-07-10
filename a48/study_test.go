package a48

import (
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct {
		name string
		n    int
	}{
		{
			name: "case1",
			n:    1,
		},
		{
			name: "case2",
			n:    2,
		},
		{
			name: "case3",
			n:    3,
		},
		{
			name: "case4",
			n:    4,
		},
		{
			name: "case5",
			n:    5,
		},
		{
			name: "case6",
			n:    6,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			f(item.n)
		})
	}
}
