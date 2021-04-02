package _202104

import "fmt"

// 168. Excel表列名称 FIXME 还未解决
func convertToTitle(columnNumber int) string {
	t := make([]byte, 0)
	for columnNumber > 26 {
		fmt.Println((columnNumber - 1) % 26)
		t = append(t, byte((columnNumber-1)%26+'A'))
		columnNumber = (columnNumber - 1) / 26
	}
	return string(t)
}
