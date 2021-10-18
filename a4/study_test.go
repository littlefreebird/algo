package a4

import "testing"

func TestF(t *testing.T)  {
	cases := []struct{
		name string
		s string
	}{
		{
			name: "case1",
			s: "abcd",
		},
		{
			name: "case2",
			s: "axyxbcd",
		},
		{
			name: "case3",
			s: "aaaaa",
		},
		{
			name: "case4",
			s: "xyxmbbbbbbnq",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%v\n", f(tt.s))
		})
	}
}
