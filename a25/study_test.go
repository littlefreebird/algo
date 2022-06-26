package a25

import (
	"fmt"
	"testing"
)

func TestNext(t *testing.T)  {
	cases := []struct{
		name string
		s string
	} {
		{
			name: "case1",
			s: "",
		},
		{
			name: "case2",
			s: "a",
		},
		{
			name: "case3",
			s: "ab",
		},
		{
			name: "case4",
			s: "aa",
		},
		{
			name: "case5",
			s: "aab",
		},
		{
			name: "case6",
			s: "abcdefg",
		},
		{
			name: "case7",
			s: "abcabefg",
		},
		{
			name: "case8",
			s: "ababaefg",
		},
		{
			name: "case9",
			s: "aaaaaaaaaag",
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			r := next(item.s)
			fmt.Println(item.s)
			fmt.Println(r)
		})
	}
}

func TestF(t *testing.T)  {
	cases := []struct{
		name string
		s string
		sub string
		pos int
	} {
		{
			name: "case1",
			s: "",
			sub: "a",
			pos: -1,
		},
		{
			name: "case2",
			s: "a",
			sub: "",
			pos: -1,
		},
		{
			name: "case3",
			s: "aaaaaa",
			sub: "aa",
			pos: 0,
		},
		{
			name: "case4",
			s: "aaaaabc",
			sub: "abc",
			pos: 4,
		},
		{
			name: "case5",
			s: "adcadcadxyz",
			sub: "adcadxyz",
			pos: 3,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			pos := kmp(item.s, item.sub)
			if pos != item.pos {
				t.Fail()
			}
		})
	}
}