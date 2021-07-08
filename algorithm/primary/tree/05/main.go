package _05

import "myleetcode/util"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = util.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var n = (len(nums) - 1) / 2
	var root = &TreeNode{Val: nums[n]}
	root.Left = sortedArrayToBST(nums[:n])
	root.Right = sortedArrayToBST(nums[n+1:])
	return root
}
