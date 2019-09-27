package _20190905

import "testing"

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	length := removeElement(nums, 3)
	if length != 2 {
		t.Fail()
	}
	for i := 0; i < length; i++ {
		t.Log(1, nums[i])
	}

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	length = removeElement(nums, 2)
	if length != 5 {
		t.Fail()
	}
	for i := 0; i < length; i++ {
		t.Log(2, nums[i])
	}
}
