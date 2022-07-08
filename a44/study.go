package a44

//给定一个非负整数数组，从第一个元素是否可达最后一个元素，每一步的大小为当前元素值大小

func f(a []int) bool {
	arrived := make([]bool, len(a))
	f1(a, 0, arrived)
	return arrived[0]
}

func f1(a []int, idx int, arrived []bool) {
	if idx == len(a)-1 {
		arrived[idx] = true
		return
	}
	f1(a, idx+1, arrived)
	if a[idx] == 0 {
		arrived[idx] = false
	} else {
		for i := 1; i <= a[idx]; i++ {
			if idx+i < len(a)-1 {
				if arrived[idx+i] {
					arrived[idx] = true
					return
				}
			} else {
				arrived[idx] = true
				return
			}
		}
		arrived[idx] = false
	}
}
