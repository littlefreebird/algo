package a45

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		s    string
		l    int
	}{
		{
			name: "case1",
			s:    " abc xyz gogo 12345 ",
			l:    5,
		},
		{
			name: "case2",
			s:    "1234",
			l:    4,
		},
		{
			name: "case3",
			s:    " 23 ",
			l:    2,
		},
		{
			name: "case4",
			s:    " 23456",
			l:    5,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			if item.l != f(item.s) {
				t.Fail()
			}
		})
	}
}
