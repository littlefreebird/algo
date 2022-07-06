package a34

import "sort"

//给定一个数组和一个target，找出数组中和为target的所有组合，且组合中不允许存在数组中相当位置的元素

func dfs(a []int, target int, pos int, p []int, ret [][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		ret = append(ret, p)
		return
	}
	for i := pos; i < len(a); i++ {
		if a[i] < target {
			if i-1 >= 0 && a[i] == a[i-1] {
				continue
			}
			p = append(p, a[i])
			dfs(a, target-a[i], i, p, ret)
			p = p[0 : len(p)-1]
		}
	}
}

func f(a []int, target int) {
	sort.Ints(a)
	var p []int
	var ret [][]int
	dfs(a, target, 0, p, ret)
}
