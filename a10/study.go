package a10

import (
	"log"
)

/*
盛最多水的容器
给出n个非负整数a1,a2,a3,a4...an，每个数代表坐标轴中的一个点(1, a1), (2, a2), ... (n, an)
在坐标内画n条垂线，垂直线i的两个端点分别是(i, 0)和(i, ai)，找出其中两条线，使得他们与x轴共同构成的
容器可以装最多的水
*/

func f(arr []int) {
	left := 0
	right := len(arr) - 1
	maxLeft := left
	maxRight := right
	max := 0
	if arr[maxLeft] < arr[maxRight] {
		max = arr[maxLeft] * (right - left + 1)
	} else {
		max = arr[maxRight] * (right - left + 1)
	}
	for {
		if right - left < 1 {
			break
		}
		if arr[left] < arr[right] {
			tmp := (right - left + 1) * arr[left]
			if tmp > max {
				max = tmp
				maxRight = right
			}
			left++
		} else {
			tmp := (right - left + 1) * arr[right]
			if tmp > max {
				max = tmp
				maxLeft = left
			}
			right--
		}
	}
	log.Printf("(%d, %d), (%d, %d)\n", maxLeft, arr[maxLeft], maxRight, arr[maxRight])
}