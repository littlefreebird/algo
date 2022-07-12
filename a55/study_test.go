package a55

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
	}{
		{
			name: "case1",
			arr:  []int{3, 312, 32},
		},
		{
			name: "case2",
			arr:  []int{3, 345, 61},
		},
		{
			name: "case3",
			arr:  []int{9, 793, 1234, 42321},
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			fmt.Println(f(item.arr))
		})
	}
}
