package a16

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		digits string
	} {
		{
			name: "case1",
			digits: "23",
		},
		{
			name: "case2",
			digits: "3456",
		},
		{
			name: "case3",
			digits: "6789",
		},
		{
			name: "case4",
			digits: "5",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%s -> %v\n", tt.digits, f(tt.digits))
		})
	}
}