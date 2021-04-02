package _201910

/* 66. 加一 */

func plusOne(digits []int) []int {
	size, n := len(digits), 1
	for i := size - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			n = 1
		} else {
			digits[i] += n
			n = 0
			break
		}
	}
	if n == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}
