package a23

import "fmt"

//删除有序数组中的重复项

func f(arr []int) int {
	fmt.Println(arr)
	pos := 0
	last := 0
	ret := 0
	for i := 0; i < len(arr); i++ {
		if i == last {
			ret++
			continue
		}
		if i > last && arr[i] == arr[last] {
			last++
			continue
		}
		if i > last && arr[i] != arr[last] {
			ret++
			arr[pos+1] = arr[i]
			pos++
			last = i
		}
	}
	fmt.Println(arr)
	return ret
}