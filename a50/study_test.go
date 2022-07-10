package a50

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct {
		name string
		m    int
		n    int
	}{
		{
			name: "case1",
			m:    2,
			n:    1,
		},
		{
			name: "case2",
			m:    1,
			n:    2,
		},
		{
			name: "case3",
			m:    3,
			n:    7,
		},
		{
			name: "case4",
			m:    3,
			n:    2,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			fmt.Println(f(item.m, item.n))
		})
	}
}
