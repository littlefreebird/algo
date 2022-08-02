package a82

import "testing"

func TestF(t *testing.T) {
	cases := []struct {
		name string
		s    string
	}{
		{
			name: "case1",
			s:    "1234",
		},
		{
			name: "case2",
			s:    "12",
		},
		{
			name: "case3",
			s:    "1",
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			f(item.s)
		})
	}
}
