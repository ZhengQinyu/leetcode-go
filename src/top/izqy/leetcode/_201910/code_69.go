package _201910

/* 69. x 的平方根 */

func mySqrt(x int) int {
	r := x
	m := r / 2
	for r/m < m {
		m = r/m + 1
	}
	return m
}
