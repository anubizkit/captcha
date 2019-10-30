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

	if pettern == 1 {
		return strconv.Itoa(left) + operstr[oper] + numstr[right]
	} else {
		return numstr[right] + operstr[oper] + strconv.Itoa(left)
	}
}
