package leetcode

import "testing"

func TestReverse(t *testing.T) {
	if x := reverse(123); x != 321 {
		t.Fail()
	}
	if x := reverse(-123); x != -321 {
		t.Fail()
	}
	if x := reverse(120); x != 21 {
		t.Fail()
	}
}
