package _20190927

/* 78. 子集 ( 未解决) */
func subsets(nums []int) [][]int {
	/* 90357 */
	var result = make([][]int, 0)
	result = append(result, make([]int, 0)) // 初始化空集合
	for _, num := range nums {
		length := len(result)
		for idx := 0; idx < length; idx++ {
			src := result[idx]
			src = append(src, num)
			result = append(result, src)
		}
	}
	return result
}
