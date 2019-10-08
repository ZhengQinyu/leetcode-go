package _201909

import (
	"testing"
)

// 整数范围 输入确保在 1 到 3999 的范围内

func TestIntToRoman(t *testing.T) {
	if intToRoman(3) != "III" {
		t.Fail()
	}
	if intToRoman(4) != "IV" {
		t.Fail()
	}
	if intToRoman(9) != "IX" {
		t.Fail()
	}
	if intToRoman(58) != "LVIII" {
		t.Fail()
	}
	if intToRoman(1994) != "MCMXCIV" {
		t.Fail()
	}
}
