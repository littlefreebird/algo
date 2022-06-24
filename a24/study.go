package a24

import "fmt"

//给定一个数组和一个元素，删除数组中同给定元素相等的项

func f(arr []int, k int) int {
	pos := 0
	end := 0
	ret := 0
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] == k {
			end++
			continue
		}
		if arr[i] != k {
			ret++
			if end > pos {
				arr[pos] = arr[i]
				pos++
				end = i
				continue
			}

		}
	}
	fmt.Println(arr)
	return ret
}