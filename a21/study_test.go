package a21

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct{
		name string
		n int
	} {
		{
			name: "case1",
			n: 0,
		},
		{
			name: "case2",
			n: 9,
		},
		{
			name: "case3",
			n: 1,
		},
		{
			name: "case4",
			n: 2,
		},
		{
			name: "case5",
			n: 3,
		},
		{
			name: "case6",
			n: 8,
		},
		{
			name: "case7",
			n: 7,
		},
	}
	for _, item := range cases{
		t.Run(item.name, func(t *testing.T) {
			fmt.Println(f(item.n))
		})
	}
}