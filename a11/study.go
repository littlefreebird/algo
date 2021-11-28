package a11

import "log"

/*
罗马数字转阿拉伯数字
I 1
V 5
X 10
L 50
C 100
D 500
M 1000
*/

func f(roman string) int {
	sum := 0
	len := len(roman)
	for i := 0; i < len; i++ {
		switch roman[i] {
		case 'I':
			if i < len - 1 {
				next := roman[i + 1]
				if next == 'V' {
					i++
					sum += 4
					continue
				} else if next == 'X' {
					i++
					sum += 9
					continue
				}
			}
			sum += 1
			continue
		case 'V':
			sum += 5
			continue
		case 'X':
			if i < len - 1 {
				next := roman[i + 1]
				if next == 'L' {
					i++
					sum += 40
					continue
				} else if next == 'C' {
					i++
					sum += 90
					continue
				 }
			}
			sum += 10
			continue
		case 'L':
			sum += 50
			continue
		case 'C':
			if i < len - 1 {
				next := roman[i + 1]
				if next == 'D' {
					i++
					sum += 400
					continue
				} else if next == 'M' {
					i++
					sum += 900
					continue
				}
			}
			sum += 100
			continue
		case 'D':
			sum += 500
			continue
		case 'M':
			sum += 1000
			continue
		default:
			log.Fatalln("unknown character")
		}
	}
	return sum
}
