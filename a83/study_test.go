package a83

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		s    string
	}{
		{
			name: "case1",
			s:    "-",
		},
		{
			name: "case2",
			s:    "123",
		},
		{
			name: "case3",
			s:    "-123",
		},
		{
			name: "case4",
			s:    "123rt",
		},
		{
			name: "case5",
			s:    "123K",
		},
		{
			name: "case5",
			s:    "204M",
		},
		{
			name: "case6",
			s:    "34G",
		},
		{
			name: "case7",
			s:    "1233.4",
		},
		{
			name: "case8",
			s:    "1233.4K",
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			f(item.s)
		})
	}
}
