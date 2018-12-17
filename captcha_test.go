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

func TestCase1112 (t *testing.T) {
	result := captcha.Captcha(1,1,1,2)
	expect := "1 + two"
	if result != expect {
		t.Error("is should be "+expect)
	}
}

func TestCase1119 (t *testing.T) {
	result := captcha.Captcha(1,1,1,9)
	expect := "1 + nine"
	if result != expect {
		t.Error("is should be "+expect)
	}
}

func TestCase1219 (t *testing.T) {
	result := captcha.Captcha(1,2,1,9)
	expect := "2 + nine"
	if result != expect {
		t.Error("is should be "+expect)
	}
}

func TestCase1919 (t *testing.T) {
	result := captcha.Captcha(1,9,1,9)
	expect := "9 + nine"
	if result != expect {
		t.Error("is should be "+expect)
	}
}

func TestCase1929 (t *testing.T) {
	result := captcha.Captcha(1,9,2,9)
	expect := "9 - nine"
	if result != expect {
		t.Error("is should be "+expect)
	}
}

func TestCase1939 (t *testing.T) {
	result := captcha.Captcha(1,9,3,9)
	expect := "9 x nine"
	if result != expect {
		t.Error("is should be "+expect)
	}
}