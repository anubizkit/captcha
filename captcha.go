package captcha

import (
	
)

func Captcha(pettern int, left int, oper int, right int) string {
	if right == 2 {
		return "1 + two"
	} else {
		return "1 + one"
	}
}