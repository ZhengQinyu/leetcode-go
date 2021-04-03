package _202104

import (
	"bytes"
)

// 168. Excel表列名称
func convertToTitle(columnNumber int) string {
	var buffer bytes.Buffer
	for columnNumber > 0 {
		columnNumber--
		buffer.WriteByte(byte('A' + (columnNumber % 26)))
		columnNumber /= 26
	}
	s := buffer.Bytes()
	length := len(s)
	half := length / 2
	for i := 0; i < half; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
	return buffer.String()
}
