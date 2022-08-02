package _202104

import "bytes"

// 482. 密钥格式化 未解决
func licenseKeyFormatting(S string, K int) string {
	var buffer bytes.Buffer
	count := 0
	for _, ch := range S {
		if ch != '-' {
			count++
		}
	}
	count %= K
	for _, ch := range S {
		if ch >= 'a' && ch <= 'z' {
			buffer.WriteByte(byte(ch - 32))
		} else {
			if ch == '-' {
				continue
			}
			buffer.WriteByte(byte(ch))
		}
	}
	//length := len(buffer.Bytes())
	return ""
}
