package a12

/*
阿拉伯数字转罗马数字
I 1
V 5
X 10
L 50
C 100
D 500
M 1000
*/

func f(d int) string {
	ret := ""
	thousand := d / 1000
	rest := d % 1000
	if thousand > 0 {
		ret += "MMM"
	}
	hundred := rest / 100
	rest = rest % 100
	if hundred > 0 {
		switch hundred {
		case 9:
			ret += "CM"
			break
		case 8:
			ret += "DCCC"
			break
		case 7:
			ret += "DCC"
			break
		case 6:
			ret += "DC"
			break
		case 5:
			ret += "D"
			break
		case 4:
			ret += "CD"
			break
		case 3:
			ret += "CCC"
			break
		case 2:
			ret += "CC"
			break
		case 1:
			ret += "C"
			break
		}
	}
	ten := rest / 10
	rest = rest % 10
	if ten > 0 {
		switch ten {
		case 9:
			ret += "IX"
			break
		case 8:
			ret += "LXXX"
			break
		case 7:
			ret += "LXX"
			break
		case 6:
			ret += "LX"
			break
		case 5:
			ret += "L"
			break
		case 4:
			ret += "XL"
			break
		case 3:
			ret += "XXX"
			break
		case 2:
			ret += "XX"
			break
		case 1:
			ret += "X"
			break
		}
	}
	if rest > 0 {
		switch rest {
		case 9:
			ret += "XC"
			break
		case 8:
			ret += "VIII"
			break
		case 7:
			ret += "VII"
			break
		case 6:
			ret += "VI"
			break
		case 5:
			ret += "V"
			break
		case 4:
			ret += "IV"
			break
		case 3:
			ret += "III"
			break
		case 2:
			ret += "II"
			break
		case 1:
			ret += "I"
			break
		}
	}
	return ret
}