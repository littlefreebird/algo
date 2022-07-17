package a79

//输入两个整数序列，第一个序列是栈的入栈序列且入栈数字都不相同，判断第二个序列是不是出栈序列

type stack struct {
	arr []int
	top int
}

func f(in []int, out []int) bool {
	var s stack
	idx := 0
	for _, oi := range out {
		for {
			if s.top != 0 && s.arr[s.top-1] == oi {
				s.top--
				break
			}
			s.arr[s.top] = in[idx]
			idx++
			s.top++
		}
	}
	return s.top == 0
}
