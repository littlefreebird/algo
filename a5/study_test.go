package a5

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		d int
	} {
		{
			name: "case1",
			d: 12345,
		},
		{
			name: "case2",
			d: -321,
		},
		{
			name: "case3",
			d: 1200,
		},
		{
			name: "case4",
			d: 1,
		},
		{
			name: "case5",
			d: 0,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%d -> %d\n", tt.d, f(tt.d))
		})
	}
}
