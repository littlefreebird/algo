package a36

//两个非负数字字符串，计算两个字符串的乘积字符串

func f(a string, b string) string {
	sum := make([]int32, len(a)+len(b))
	diff := 1 - '1'
	for j, bi := range b {
		for i, ai := range a {
			sum[i+j+1] = (ai+diff)*(bi+diff) + sum[i+j+1]
		}
	}
	for k := len(a) + len(b) - 1; k >= 0; k-- {
		tmp := sum[k]
		sum[k] = tmp % 10
		if k > 0 {
			sum[k-1] += tmp / 10
		}
	}
	start := true
	var ret string
	for i, item := range sum {
		if i == 0 && item == 0 {
			start = false
		}
		if item != 0 {
			start = true
		}
		if start {
			ret += string(item - diff)
		}
	}
	if ret == "" {
		ret = "0"
	}
	return ret
}
