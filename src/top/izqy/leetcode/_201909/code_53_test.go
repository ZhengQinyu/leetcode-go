package _201909

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	if x := maxSubArray(nums); x != 6 {
		t.Fail()
	}

	nums = []int{-2}
	if x := maxSubArray(nums); x != -2 {
		t.Fail()
	}
}
