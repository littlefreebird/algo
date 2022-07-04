package a26

import "fmt"

//给定两个整数求其整数商，不能使用乘法除法和mod运算

func f1(a, b int) int {
	lc := -1
	last := b
	x := a
	y := b
	for {
		fmt.Println(x, y)
		if x-y == 0 {
			return 1 << (lc + 1)
		} else if x-y < 0 {
			if lc < 0 {
				return 0
			} else {
				fmt.Println(lc)
				return 1<<lc + f1(a-last, b)
			}
		} else {
			lc++
			last = y
			y += y
		}
	}
}

func f(a, b int) int {
	if b == 0 {
		return -1
	}
	if a == 0 {
		return 0
	}
	x := a
	y := b
	same := false
	if x > 0 && y > 0 || x < 0 && y < 0 {
		same = true
	}
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	r := f1(x, y)
	if !same {
		return -r
	}
	return r
}
