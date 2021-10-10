/*
给定整数数组nums和一个整数目标值target，在数组中找出和为target的两个数的下标
输入：nums = [2, 7, 11, 15], target = 9
输出：[0, 1]
*/

package main

import "fmt"

func f(nums []int, target int) {
	num2idx := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num2idx[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if idx, ok := num2idx[target - nums[i]]; ok {
			if idx <= i {
				continue
			}
			fmt.Printf("%v\n", []int{i, idx})
		}
	}
}


