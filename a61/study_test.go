package a61

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		s    string
		p    string
		r    bool
	}{
		{
			name: "case1",
			p:    ".a*bcd",
			s:    "abcd",
			r:    true,
		},
		{
			name: "case2",
			p:    "abc.defg*x",
			s:    "abcddefgxxxx",
			r:    false,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			if item.r != f(item.s, item.p) {
				t.Fail()
			}
		})
	}
}
