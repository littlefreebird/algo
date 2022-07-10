package a47

import (
	"fmt"
)

//给定一个无重叠的区间数组，数组按照区间的start排序，在数组中插入一个区间，确保数组中区间依然有序且不重叠

type section struct {
	start int
	end   int
}

func f(intervals []section, x section) {
	fmt.Println(intervals)
	fmt.Println(x)
	var ret []section
	start := x.start
	end := x.end
	idx := -1
	for _, item := range intervals {
		idx++
		if item.start > end {
			ret = append(ret, section{start: start, end: end})
			break
		}
		if item.start >= start && end >= item.start && end <= item.end {
			end = item.end
			ret = append(ret, section{start: start, end: end})
			idx++
			break
		}
		if item.start >= start && item.end < end {
			if idx == len(intervals)-1 {
				ret = append(ret, section{start: start, end: end})
				idx++
			}
			continue
		}
		if start > item.end {
			ret = append(ret, item)
			if idx == len(intervals)-1 {
				ret = append(ret, section{start: start, end: end})
				idx++
			}
			continue
		}
		if start >= item.start && start <= item.end {
			start = item.start
			if idx == len(intervals)-1 {
				ret = append(ret, section{start: start, end: end})
				idx++
			}
			continue
		}
	}
	if idx < len(intervals) {
		ret = append(ret, intervals[idx:]...)
	}
	fmt.Println(ret)
}
