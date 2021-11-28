package a13

import "testing"

func TestF(t *testing.T) {
	cases := []struct{
		name string
		array []string
	} {
		{
			name: "case1",
			array: []string{"", ""},
		},
		{
			name: "case2",
			array: []string{"", "abc"},
		},
		{
			name: "case3",
			array: []string{"abc", "abd", "abg"},
		},
		{
			name: "case4",
			array: []string{"abcdefg", "ghi", "xyz"},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%v\n%s", tt.array, f(tt.array))
		})
	}
}
