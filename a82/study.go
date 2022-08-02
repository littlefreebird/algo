package a82

//列出一个字符串的所有子集

import "fmt"

func f(s string) {
	var set []string
	for i := 0; i < len(s); i++ {
		f1(s, i, "", &set)
	}
	fmt.Println(s)
	fmt.Println(set)
}

func f1(s string, end int, current string, set *[]string) {
	if end >= len(s) {
		return
	}
	*set = append(*set, current+string(s[end]))
	for i := end + 1; i < len(s); i++ {
		f1(s, i, current+string(s[end]), set)
	}
}
