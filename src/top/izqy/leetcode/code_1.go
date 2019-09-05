package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		num := target - v
		if idx, ok := m[num]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
