package a15

import (
	"log"
)

/*给定一个整数数组和一个target，从数组中找出三个数，其和与target最接近*/

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

func f(array []int, target int) {
	if len(array) < 3 {
		log.Fatalln("array length is less than 3")
	}
	sort(array)
	tl := 1
	tr := len(array) - 1
	ti := 0
	for i := 0; i < len(array); i++ {
		left := i + 1
		right := len(array) - 1

		for {
			if left >= right {
				break
			}
			sum := array[i] + array[left] + array[right]
			if sum == target {
				log.Println(array[i], array[left], array[right])
				return
			} else if sum < target {
				do := array[ti] + array[tl] + array[tr] - sum
				if do < 0 {
					do = -1 * do
				}
				dn := array[i] + array[left] + array[right] - sum
				if dn < 0 {
					dn = -1 * dn
				}
				if dn < do {
					tl = left
					tr = right
					ti = i
				}
				left++
			} else {
				do := array[ti] + array[tl] + array[tr] - sum
				if do < 0 {
					do = -1 * do
				}
				dn := array[i] + array[left] + array[right] - sum
				if dn < 0 {
					dn = -1 * dn
				}
				if dn < do {
					tl = left
					tr = right
					ti = i
				}
				right--
			}
		}
	}
	log.Println(array[ti], array[tl], array[tr])
}