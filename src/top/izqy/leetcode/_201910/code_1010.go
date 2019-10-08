package _201910

import "fmt"

/* 1010. 总持续时间可被 60 整除的歌曲 未解决 */
func numPairsDivisibleBy60(time []int) int {
	dir := make([]int, 60)
	for _, t := range time {
		dir[t%60]++
	}
	count := 0
	for idx, t := range dir {
		fmt.Println(idx)
		m := t % 60
		if dir[m] > 0 {
			if m == 0 || m == 30 {
				count += (dir[m] * (dir[m] - 1)) / 2
			} else {
				count += dir[m] * dir[60-m]
				dir[60-m] = 0 // 重置为0避免重复统计
			}
		}
		if idx >= 30 {
			break
		}
	}
	return count
}
