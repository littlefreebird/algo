package a46

import "fmt"

/*以数组intervals表示若干个区间的集合，其中单个区间为intervals[i] = [start, end]，请合并所有重叠区间，
返回一个不重叠的区间数组*/

type section struct {
	start int
	end   int
}

func quickSort(intervals []section, l int, r int) {
	if l < r {
		x := intervals[l]
		i := l
		j := r
		for {
			if i >= j {
				break
			}
			for {
				if j > i {
					if intervals[j].start > x.start || (intervals[j].start == x.start && intervals[j].end > x.end) {
						j--
					} else {
						intervals[i] = intervals[j]
						i++
						break
					}
				} else {
					break
				}
			}
			for {
				if j > i {
					if intervals[i].start < x.start || (intervals[i].start == x.start && intervals[i].end < x.end) {
						i++
					} else {
						intervals[j] = intervals[i]
						j--
						break
					}
				} else {
					break
				}
			}
		}
		intervals[i] = x
		quickSort(intervals, l, i-1)
		quickSort(intervals, i+1, r)
	}
}

func f(intervals []section) {
	quickSort(intervals, 0, len(intervals)-1)
	fmt.Println(intervals)
	i := intervals[0].start
	j := intervals[0].end
	var ret []section
	for idx, item := range intervals {
		if item.start > j {
			ret = append(ret, section{
				start: i,
				end:   j,
			})
			i = item.start
			j = item.end
		} else {
			if item.end > j {
				j = item.end
			}
		}
		if idx == len(intervals)-1 && len(intervals) > 1 {
			ret = append(ret, section{
				start: i,
				end:   j,
			})
		}
	}
	fmt.Println(ret)
}
