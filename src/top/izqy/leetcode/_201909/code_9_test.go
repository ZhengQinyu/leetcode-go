package _201909

import "testing"

func TestIsPalindrome(t *testing.T) {
	if isPalindrome(121) {
		t.Log("ok")
	}
	if isPalindrome(-121) {
		t.Fail()
	}
	if isPalindrome(10) {
		t.Fail()
	}
}
