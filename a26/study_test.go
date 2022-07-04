package a26

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct {
		name string
		a    int
		b    int
		r    int
	}{
		{
			name: "case1",
			a:    0,
			b:    1,
			r:    0,
		},
		{
			name: "case2",
			a:    1,
			b:    0,
			r:    -1,
		},
		{
			name: "case3",
			a:    13,
			b:    2,
			r:    6,
		},
		{
			name: "case4",
			a:    16,
			b:    4,
			r:    4,
		},
		{
			name: "case5",
			a:    -14,
			b:    -3,
			r:    4,
		},
		{
			name: "case6",
			a:    -15,
			b:    5,
			r:    -3,
		},
		{
			name: "case7",
			a:    -15,
			b:    6,
			r:    -2,
		},
		{
			name: "case8",
			a:    20,
			b:    -5,
			r:    -4,
		},
		{
			name: "case9",
			a:    20,
			b:    -7,
			r:    -2,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			fmt.Println(item)
			if item.r != f(item.a, item.b) {
				t.Fail()
			}
		})
	}
}
