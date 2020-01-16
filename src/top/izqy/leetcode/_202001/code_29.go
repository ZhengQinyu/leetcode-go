package _202001

func divide(dividend int, divisor int) int {
	count := 0
	for dividend >= divisor {
		divisor += divisor // 翻倍
		if count == 0 {
			count = 1
		} else {
			count += count
		}
	}
	return count
}
