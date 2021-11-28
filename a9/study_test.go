package a9

import (
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct{
		name string
		d int
	}{
		{
			name: "case1",
			d: 0,
		},
		{
			name: "case2",
			d: 9,
		},
		{
			name: "case3",
			d: 10,
		},
		{
			name: "case4",
			d: 4455,
		},
		{
			name: "case5",
			d: 101,
		},
		{
			name: "case6",
			d: 9009,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%d -> %t\n", tt.d, f(tt.d))
		})
	}
}
