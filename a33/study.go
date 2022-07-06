package a33

//给定一个无重复数字的数组和一个target，在数组中找出所有和为target的数字组合，组合中数字可重复

func f(a []int, target int) [][]int {
	var p []int
	var ret [][]int
	dfs(a, target, 0, p, ret)
	return ret
}

func dfs(a []int, target int, idx int, p []int, paths [][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		paths = append(paths, p)
		return
	}
	for i := idx; i < len(a); i++ {
		if a[i] < target {
			p = append(p, a[i])
			dfs(a, target-a[i], i, p, paths)
			p = p[0 : len(p)-1]
		}
	}
}
