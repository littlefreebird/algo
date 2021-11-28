package a14

import "log"

/*一个整数数组，找出其中和为0的三个数字且不重复*/

func sort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := len(array) - 1; j > i; j-- {
			if array[j] < array[j - 1] {
				tmp := array[j]
				array[j] = array[j - 1]
				array[j - 1] = tmp
			}
		}
	}
}

func f(array []int) {
	var res [][]int
	sort(array)
	for i := 0; i < len(array); i++ {
		if array[i] > 0 {
			break
		}
		left := i + 1
		right := len(array) - 1
		for {
			if left >= right {
				break
			}
			sum := array[i] + array[left] + array[right]
			if sum == 0 {
				r := []int{array[i], array[left], array[right]}
				res = append(res, r)
				for {
					if right - left > 1 && array[left] == array[left + 1] {
						left++
					} else {
						break
					}
				}
				left++
				for {
					if right - left > 1 && array[right] == array[right - 1] {
						right--
					} else {
						break
					}
				}
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	log.Printf("%v\n", res)
}