package captcha

import (
	"strconv"
)

func Captcha(pettern int, left int, oper int, right int) string {
	return strconv.Itoa(left) + IntOperToStr(oper) + IntToStr(right)
}

func IntToStr(num int) string {
	switch num {
		case 0 : return "zero"
		case 1 : return "one"
		case 2 : return "two"
		case 3 : return "three"
		case 4 : return "four"
		case 5 : return "five"
		case 6 : return "six"
		case 7 : return "seven"
		case 8 : return "eight"
		case 9 : return "nine"
	}
	return ""
}

func IntOperToStr(num int) string {
	switch num {
		case 1 : return " + "
		case 2 : return " - "
		case 3 : return " x "
	}
	return ""
}