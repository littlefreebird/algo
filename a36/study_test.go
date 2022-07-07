package a36

import (
	"testing"
)

func TestF(t *testing.T) {
	cases := []struct {
		name string
		a    string
		b    string
		r    string
	}{
		{
			name: "case1",
			a:    "6789",
			b:    "98",
			r:    "665322",
		},
		{
			name: "case2",
			a:    "60895",
			b:    "403",
			r:    "24540685",
		},
		{
			name: "case3",
			a:    "0",
			b:    "12345",
			r:    "0",
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			if item.r != f(item.a, item.b) {
				t.Fail()
			}
		})
	}
}
