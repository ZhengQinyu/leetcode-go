package _20190905

import "testing"

func TestIsValid(t *testing.T) {
	if !isValid("()") {
		t.Fail()
	}
	if !isValid("()[]{}") {
		t.Fail()
	}
	if isValid("(]") {
		t.Fail()
	}
	if isValid("([)]") {
		t.Fail()
	}
	if !isValid("{[]}") {
		t.Fail()
	}
}
