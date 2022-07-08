package a42

import "fmt"

//给定一个整数数组，找出数组中和最大子数组

type section struct {
	start int
	end   int
}

type maxInfo struct {
	start   int
	end     int
	current int
	max     int
}

func f(a []int) {
	subMax := make(map[section]maxInfo)
	f1(a, 0, len(a)-1, subMax)
	fmt.Println(subMax[section{start: 0, end: len(a) - 1}])
}

func f1(a []int, start int, end int, subMax map[section]maxInfo) {
	if end == start {
		subMax[section{start: start, end: end}] = maxInfo{start: start, end: end, current: a[start], max: a[start]}
		return
	}
	f1(a, start+1, end, subMax)
	f1(a, start, end-1, subMax)
	var max1 int
	var max2 int
	var maxInfo1 maxInfo
	var maxInfo2 maxInfo
	maxLeftInfo := subMax[section{start: start, end: end - 1}]
	maxRightInfo := subMax[section{start: start + 1, end: end}]
	if a[start]+maxRightInfo.current > maxRightInfo.max {
		max1 = a[start] + maxRightInfo.current
		maxInfo1.start = start
		maxInfo1.end = end
		maxInfo1.current = max1
		maxInfo1.max = max1
	} else {
		max1 = maxRightInfo.max
		maxInfo1 = maxRightInfo
		maxInfo1.current = maxInfo1.current + a[start]
	}
	if a[end]+maxLeftInfo.current > maxLeftInfo.max {
		max2 = a[end] + maxLeftInfo.current
		maxInfo2.start = start
		maxInfo2.end = end
		maxInfo2.current = max2
		maxInfo2.max = max2
	} else {
		max2 = maxLeftInfo.max
		maxInfo2 = maxLeftInfo
		maxInfo2.current = a[end] + maxInfo2.current
	}
	if max1 > max2 {
		subMax[section{start: start, end: end}] = maxInfo1
	} else {
		subMax[section{start: start, end: end}] = maxInfo2
	}
}
