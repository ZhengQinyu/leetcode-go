package leetcode

import "testing"

func TestStrStr(t *testing.T) {
	if strStr("hello", "ll") != 2 {
		t.Fail()
	}
	if strStr("aaaaa", "bba") != -1 {
		t.Fail()
	}
	if strStr("aaaaa", "") != 0 {
		t.Fail()
	}
	if strStr("abcdabcdefg", "abcde") != 4 {
		t.Fail()
	}
}
