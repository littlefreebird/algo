package a11

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		roman string
	} {
		{
			name: "case1",
			roman: "",
		},
		{
			name: "case2",
			roman: "IV",
		},
		{
			name: "case3",
			roman: "LVIII",
		},
		{
			name: "case4",
			roman: "MDCLXVI",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%s -> %d\n", tt.roman, f(tt.roman))
		})
	}
}
