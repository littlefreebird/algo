package a12

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		d int
	} {
		{
			name: "case1",
			d: 3675,
		},
		{
			name: "case2",
			d: 0,
		},
		{
			name: "case3",
			d: 1234,
		},
		{
			name: "case4",
			d: 908,
		},
		{
			name: "case5",
			d: 30,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%d -> %s\n", tt.d, f(tt.d))
		})
	}
}
