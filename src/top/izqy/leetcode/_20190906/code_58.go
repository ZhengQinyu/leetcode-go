package leetcode

func lengthOfLastWord(s string) int {
	count := 0
	for _, c := range s {
		if c != ' ' {
			count++
		} else {
			return count
		}
	}
	return count
}
