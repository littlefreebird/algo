package a84

import "fmt"

//返回一个整数数组的前k个数字

func f(arr []int, k int) {
	fmt.Println(arr)
	if len(arr) == 0 || k <= 0 || k > len(arr) {
		return
	}
	i := 0
	j := len(arr) - 1
	pos := -1
	for {
		if i <= j && pos != k-1 {
			pos = f1(arr, i, j)
			if pos > k-1 {
				j = pos - 1
			}
			if pos < k-1 {
				i = pos + 1
			}
		} else {
			break
		}
	}
	fmt.Println(arr)
}

func f1(arr []int, i int, j int) int {
	t := arr[i]
	for {
		if i >= j {
			break
		}
		for {
			if j <= i {
				break
			}
			if arr[j] < t {
				arr[i] = arr[j]
				i++
				break
			}
			j--
		}
		for {
			if j <= i {
				break
			}
			if arr[i] > t {
				arr[j] = arr[i]
				j--
				break
			}
			i++
		}
	}
	arr[i] = t
	return i
}
