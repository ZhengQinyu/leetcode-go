package _20190905

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	length := removeDuplicates(nums)
	if length != 2 {
		t.Fail()
	}
	for i := 0; i < length; i++ {
		t.Log(1, nums[i])
	}

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length = removeDuplicates(nums)
	if length != 5 {
		t.Fail()
	}
	for i := 0; i < length; i++ {
		t.Log(2, nums[i])
	}
}
