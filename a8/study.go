package a8

//字符串转整数

func f(s string) int {
	if s == "" {
		return 0
	}
	start := false
	flag := true
	sum := 0
	for i := 0; i < len(s); i++ {
		if start == false {
			if s[i] == ' ' {

			} else if s[i] == '+' || s[i] == '-' || (s[i] >= '0' && s[i] <= '9') {
				start = true
				if s[i] == '-' {
					flag = false
				}
				if s[i] >= '0' && s[i] <= '9' {
					sum += int(s[i]) - '0'
				}
			} else {
				return 0
			}
		} else {
			if s[i] >= '0' && s[i] <= '9' {
				sum = sum * 10 + int(s[i]) - '0'
			} else {
				break
			}
		}
	}
	if !flag {
		sum = sum * -1
	}
	return sum
}