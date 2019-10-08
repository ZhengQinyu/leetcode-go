package _201909

import "testing"

// 所有输入只包含小写字母 a-z

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	if longestCommonPrefix(strs) != "fl" {
		t.Fail()
	}

	strs = []string{"dog", "racecar", "car"}
	if longestCommonPrefix(strs) != "" {
		t.Fail()
	}
}
