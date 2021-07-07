package _02

import (
	"math"
	"myleetcode/util"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = util.TreeNode

func isValidBST(root *TreeNode) bool {
	return isValidBSTs(root, math.MinInt64, math.MaxInt64)

}

func isValidBSTs(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}
	return isValidBSTs(root.Left, min, root.Val) && isValidBSTs(root.Right, root.Val, max)
}

func treeToInorder(node *TreeNode) []int {
	var res []int
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}
	res = treeToInorder(node.Left)
	res = append(res, node.Val)
	res = append(res, treeToInorder(node.Right)...)
	return res
}
