package _20190930

/* 190. 颠倒二进制位 */
func reverseBits(num uint32) uint32 {
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		result <<= 1
		result = result | (num & 1)
		num >>= 1
	}
	return result
}