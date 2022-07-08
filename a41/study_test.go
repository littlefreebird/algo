package a41

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct {
		name string
		b    float64
		n    int
	}{
		{
			name: "case1",
			b:    2,
			n:    1,
		},
		{
			name: "case2",
			b:    2,
			n:    2,
		},
		{
			name: "case3",
			b:    2,
			n:    3,
		},
		{
			name: "case4",
			b:    2,
			n:    4,
		},
		{
			name: "case5",
			b:    2,
			n:    5,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			r := f(item.b, item.n)
			fmt.Println(r)
		})
	}
}
