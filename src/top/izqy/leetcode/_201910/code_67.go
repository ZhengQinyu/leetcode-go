package _201910

import "bytes"

/* 67. 二进制求和 */

// 转数字的方式，字符串超出
/*func addBinary(a string, b string) string {
	// 转数字
	var n1, n2 int64
	n1, l1 := 0, len(a)
	for i := 0; i < l1; i++ {
		n1 <<= 1
		if a[i] == '1' {
			n1 |= 1
		}
	}
	n2, l2 := 0, len(b)
	for i := 0; i < l2; i++ {
		n2 <<= 1
		if b[i] == '1' {
			n2 |= 1
		}
	}
	return strconv.FormatInt(n1+n2, 2)
}*/

func addBinary(a string, b string) string {
	l1, l2 := len(a), len(b)
	var result []byte
	var sum, n byte
	for x := 0; x < l1 || x < l2; x++ {
		if x >= l1 {
			sum = b[l2-1-x] - '0' + n
		} else if x >= l2 {
			sum = a[l1-1-x] - '0' + n
		} else {
			sum = a[l1-1-x] - '0' + b[l2-1-x] - '0' + n
		}
		if sum >= 2 {
			n = 1
			result = append(result, sum%2)
		} else {
			result = append(result, byte(sum))
			n = 0
		}
	}
	if n == 1 {
		result = append(result, 1)
	}
	var buffer bytes.Buffer
	for i := len(result) - 1; i >= 0; i-- {
		buffer.WriteByte(result[i] + '0')
	}
	return buffer.String()
}
