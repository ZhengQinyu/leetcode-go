package leetcode

import "testing"

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	if searchInsert(nums, 5) != 2 {
		t.Fail()
	}
	if searchInsert(nums, 2) != 1 {
		t.Fail()
	}
	if searchInsert(nums, 7) != 4 {
		t.Fail()
	}
	if searchInsert(nums, 0) != 0 {
		t.Fail()
	}
}
