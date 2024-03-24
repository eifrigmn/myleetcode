package main

import "sort"

func hIndex(citations []int) int {
	if len(citations) == 1 && citations[0] > 0 {
		return 1
	}
	// 引用x次的论文数量
	var mp = make(map[int]int)
	var maxTms int
	for _, tm := range citations {
		if tm == 0 {
			continue
		}
		if tm > maxTms {
			maxTms = tm
		}
		incrTms(mp, tm)
	}
	for i := maxTms; i > 0; i-- {
		if mp[i] >= i {
			return i
		}
	}
	return 0
}

func incrTms(mp map[int]int, tm int) {
	for i := 1; i <= tm; i++ {
		mp[i]++
	}
}

func _hIndex(citations []int) int {
	sort.Ints(citations)
	h := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] > h {
			h++
		}
	}
	return h
}
