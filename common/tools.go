package common

import "bytes"

/* 辗转相除法： 计算最大公约数 */
func Gcd(m, n int) int {
	for n > 0 {
		m, n = n, m%n
	}
	return m
}

// 字符翻转
func Reverse(buffer *bytes.Buffer) string {
	s := buffer.Bytes()
	length := len(s)
	half := length / 2
	for i := 0; i < half; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
	return buffer.String()
}
