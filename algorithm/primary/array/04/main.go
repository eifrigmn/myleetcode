package _04

func containsDuplicate(nums []int) bool {
	var mp = make(map[int]int, len(nums))
	for idx, num := range nums {
		if _, ok := mp[num];ok {
			return true
		}
		mp[num] = idx
	}
	return false
}
