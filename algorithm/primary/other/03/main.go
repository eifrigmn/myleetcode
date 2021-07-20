package _03

func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		if num>>i&1 == 1 {
			result += 1 << (31 - i)
		}
	}
	return result
}
