package captcha

import (
	"strconv"
)

func Captcha(pettern int, left int, oper int, right int) string {

	numstr := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}
	operstr := map[int]string{
		1: " + ",
		2: " - ",
		3: " x ",
	}
	patternFunc := map[int]func(int, int, int) string{
		0: func(left int, oper int, right int) string { return numstr[right] + operstr[oper] + strconv.Itoa(left) },
		1: func(left int, oper int, right int) string { return strconv.Itoa(left) + operstr[oper] + numstr[right] },
	}

	return patternFunc[pettern](left, oper, right)
}
