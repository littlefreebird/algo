package a30

//给定一个非递减数组和一个 target，找出target在数组的开始位置和结束位置

func left(a []int, target int) int {
	if len(a) == 0 {
		return -1
	}
	min := 0
	max := len(a) - 1
	for {
		mid := (min + max) / 2
		if target > a[mid] {
			min = mid + 1
		} else {
			max = mid
		}
		if min >= max {
			if target == a[min] {
				return min
			} else {
				return -1
			}
		}
	}
}

func right(a []int, target int) int {
	if len(a) == 0 {
		return -1
	}
	min := 0
	max := len(a) - 1
	for {
		mid := (min + max + 1) / 2
		if target < a[mid] {
			max = mid - 1
		} else {
			min = mid
		}
		if min >= max {
			if target == a[max] {
				return max
			} else {
				return -1
			}
		}
	}
}

func f(a []int, target int) (int, int) {
	return left(a, target), right(a, target)
}
