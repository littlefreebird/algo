package a28

//给定旋转升序数组，和一个整数target，在数组中查询target输出index，如果没有找到输出-1，时间复杂度lgn

func f(a []int, target int) int {
	min := 0
	max := len(a) - 1
	for {
		mid := (min + max) / 2
		if a[mid] == target {
			return mid
		}
		if mid > max || mid < min {
			return -1
		}
		if a[mid] < a[max] {
			if target >= a[mid] && target <= a[max] {
				min = mid + 1
			} else {
				max = mid - 1
			}
		} else {
			if target >= a[min] && target <= a[mid] {
				max = mid - 1
			} else {
				min = mid + 1
			}
		}
	}
}
