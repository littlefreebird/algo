package a29

import "fmt"

//给定一个有序数组一个target，在数组中找出target的位置，如果target不存在则找出插入位置

func f(a []int, target int) int {
	if len(a) == 0 {
		return 0
	}
	incr := false
	if a[len(a)-1] >= a[0] {
		incr = true
	}
	min := 0
	max := len(a) - 1
	for {
		mid := (min + max) / 2
		if a[mid] == target {
			return mid
		}
		fmt.Println(min, max)
		if min >= max {
			if incr {
				if target > a[max] {
					return max + 1
				}
				if target < a[min] {
					return min
				}
			} else {
				if target > a[min] {
					return min
				}
				if target < a[max] {
					return max + 1
				}
			}
		}
		if incr {
			if target < a[mid] {
				max = mid - 1
			}
			if target > a[mid] {
				min = mid + 1
			}
		} else {
			if target > a[mid] {
				max = mid - 1
			}
			if target < a[mid] {
				min = mid + 1
			}
		}
	}
}
