package a5

/*
反转一个整数
输入：123
输出：321

输入：-345
输出：-543

输入：120
输出：21
*/

func f(d int) int {
	sum := 0
	mul := 10
	rest := d
	for {
		t := rest % 10
		rest = rest / 10
		sum = sum * mul + t
		if rest == 0 {
			break
		}
	}
	if d > 0 {
		return sum
	} else {
		return sum * -1
	}
}