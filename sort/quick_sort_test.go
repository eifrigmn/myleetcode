package sort

import (
	"myleetcode/util"
	"testing"
)

func TestQuickSortPtail(t *testing.T) {
	t.Log("选择尾元素为分区点")
	nums := util.GenerateSlice(5)
	result := quickSort(nums)
	t.Log(result)
}
