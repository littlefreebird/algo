package test1

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct {
		name string
		s    string
	}{
		{
			name: "case1",
			s:    "123  456",
		},
		{
			name: "case2",
			s:    "123   ",
		},
		{
			name: "case3",
			s:    "   1234",
		},
		{
			name: "case4",
			s:    "123 45    789   ",
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			fmt.Println(item.s)
			fmt.Println(f(item.s))
		})
	}
}
