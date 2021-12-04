package a17

import "log"

/*给定一个整数数组和一个target，找出数组中四个数其和等于target*/

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
	if len(array) < 4 {
		log.Fatalln("array length is les than 4")
	}
	sort(array)
	log.Printf("%v, %d\n", array, target)
	for i := 0; i < len(array); i++ {
		for j:= i + 1; j < len(array); j++ {
			left := j + 1
			right := len(array) - 1
			for  {
				if left >= right {
					break
				}
				sum := array[i] + array[j] + array[left] + array[right]
				if sum == target {
					log.Println(array[i], array[j], array[left], array[right])
					for {
						if left + 1 >= right {
							break
						}
						if array[left + 1] == array[left] {
							left++
						} else {
							break
						}
					}
					left++
					for {
						if left >= right - 1 {
							break
						}
						if array[right] == array[right - 1] {
							right--
						} else {
							break
						}
					}
					right--
				}
				if sum > target {
					right--
				}
				if sum < target {
					left++
				}
			}
		}
	}
}
