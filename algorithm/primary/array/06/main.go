package _06

func intersect(nums1 []int, nums2 []int) []int {
	var result []int
	for _, n1 := range nums1 {
		for i:=0;i<len(nums2);i++{
			if n1 == nums2[i] {
				result = append(result, n1)
				if i != len(nums2)-1 {
					nums2 = append(nums2[:i], nums2[i+1:]...)
				} else {
					nums2 = nums2[:i]
				}
				break
			}
		}
	}
	return result
}

func intersect1(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _,v:= range nums1 {
		m[v] = m[v]+1
	}
	output := []int {}
	for _,v := range nums2 {
		if m[v] >0 {
			output = append(output,v)
			m[v] = m[v] - 1
		}
	}
	return output
}
