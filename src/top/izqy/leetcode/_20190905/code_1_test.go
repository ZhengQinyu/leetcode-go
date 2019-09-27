package _20190905

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 19}
	if x := twoSum(nums, 9); x[0] != 0 || x[1] != 1 {
		t.Fail()
	}
}
