package a20

import (
	"github.com/littlefreebird/algo/common"
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct{
		name string
		l1 []int
		l2 []int
	} {
		{
			name: "case1",
			l1: []int{},
			l2: []int{},
		},
		{
			name: "case2",
			l1: []int{1, 2, 3},
			l2: []int{},
		},
		{
			name: "case3",
			l1: []int{},
			l2: []int{4, 5, 6},
		},
		{
			name: "case4",
			l1: []int{1, 3, 4, 6, 8, 9},
			l2: []int{1, 5, 7, 11, 18},
		},
	}
	for _, tt := range cases {
		common.PrintLinkList(f(tt.l1, tt.l2))
	}
}