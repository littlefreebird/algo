package a40

//给定一个字符串数组，将字母异位词组合在一起

func f(aStr []string) {
	key2Arr := make(map[int64][]string)
	for _, item := range aStr {
		key := int64(0)
		for _, i := range item {
			if i >= 'a' && i <= 'z' {
				key = key & (1 << (i - 'a'))
			}
			if i >= 'A' && i <= 'Z' {
				key = key & (1 << (26 + i - 'A'))
			}
		}
		key2Arr[key] = append(key2Arr[key], item)
	}
}
