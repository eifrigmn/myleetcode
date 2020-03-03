package sort 

func intersect1(nums1 []int, nums2 []int) []int{
	if len(nums2) > len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	var mp = make(map[int]int)
	var result []int 
	for i :=0;i<len(nums1);i++{
		if v, ok := mp[nums1[i]]; ok != false {
			mp[nums1[i]] = v+1
		} else {
			mp[nums1[i]] = 1
		}
	}

	for j:=0;j<len(nums2); j++{
		val, ok := mp[nums2[j]]
		if ok && val > 0{
			result = append(result, nums2[j])
			mp[nums2[j]] = val - 1
		}
	}

	return result 
}