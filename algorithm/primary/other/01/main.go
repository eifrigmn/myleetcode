package _01

// num每次右移1位，然后每次判断最后一位是否为1
func hammingWeight(num uint32) int {
	var count int
	for i := 0; i < 32; i++ {
		if num>>i&1 == 1 {
			count++
		}
	}
	return count
}
