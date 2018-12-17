package captcha_test

import (
	"testing"
	captcha "github.com/anubizkit/captcha"
)

func TestCase1111 (t *testing.T) {
	result := captcha.Captcha(1,1,1,1)
	expect := "1 + one"
	if result != expect {
		t.Error("is should be "+expect)
	}
}