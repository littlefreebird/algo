package a30

import "testing"

func TestLeft(t *testing.T) {
	cases := []struct {
		name   string
		arr    []int
		target int
		pos    int
	}{
		{
			name:   "case1",
			arr:    []int{1},
			target: 1,
			pos:    0,
		},
		{
			name:   "case2",
			arr:    []int{1, 2, 2, 2, 3, 4, 5},
			target: 2,
			pos:    1,
		},
		{
			name:   "case3",
			arr:    []int{1, 1, 2, 3, 4, 5},
			target: 1,
			pos:    0,
		},
		{
			name:   "case4",
			arr:    []int{1, 3, 4, 5, 5, 5, 5},
			target: 5,
			pos:    3,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			if item.pos != left(item.arr, item.target) {
				t.Fail()
			}
		})
	}
}

func TestRight(t *testing.T) {
	cases := []struct {
		name   string
		arr    []int
		target int
		pos    int
	}{
		{
			name:   "case1",
			arr:    []int{2, 2},
			target: 2,
			pos:    1,
		},
		{
			name:   "case2",
			arr:    []int{3},
			target: 3,
			pos:    0,
		},
		{
			name:   "case3",
			arr:    []int{1, 3, 4, 6, 9, 9, 9, 9},
			target: 9,
			pos:    7,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			if item.pos != right(item.arr, item.target) {
				t.Fail()
			}
		})
	}
}

func TestF(t *testing.T) {
	cases := []struct {
		name   string
		arr    []int
		target int
		left   int
		right  int
	}{
		{
			name:   "case1",
			arr:    []int{3},
			target: 3,
			left:   0,
			right:  0,
		},
		{
			name:   "case2",
			arr:    []int{3, 4, 5, 5},
			target: 1,
			left:   -1,
			right:  -1,
		},
		{
			name:   "case3",
			arr:    []int{3, 4, 5, 5, 5, 6, 7},
			target: 5,
			left:   2,
			right:  4,
		},
		{
			name:   "case4",
			arr:    []int{3, 3, 3, 6, 7, 8},
			target: 3,
			left:   0,
			right:  2,
		},
		{
			name:   "case5",
			arr:    []int{4, 4, 4},
			target: 4,
			left:   0,
			right:  2,
		},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			l, r := f(item.arr, item.target)
			if l != item.left || r != item.right {
				t.Fail()
			}
		})
	}
}
