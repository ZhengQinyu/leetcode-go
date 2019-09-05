package leetcode

func searchInsert(nums []int, target int) int {
	length := len(nums)
	i := 0
	for i < length && nums[i] < target {
		i++
	}
	return i
}
