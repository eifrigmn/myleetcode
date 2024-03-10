package _0160

func twoSum(numbers []int, target int) []int {
	var lft, rgt = 0, len(numbers) - 1
	for lft < rgt {
		if numbers[lft]+numbers[rgt] > target {
			rgt--
		} else if numbers[lft]+numbers[rgt] < target {
			lft++
		} else {
			return []int{lft + 1, rgt + 1}
		}
	}
	return nil
}
