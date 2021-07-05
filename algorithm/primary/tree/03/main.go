package _03

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

func isSymmetric(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}
	var q = []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			if q[i] == nil && q[size-i-1] == nil {
				continue
			}
			if q[i] != nil && q[size-i-1] != nil && q[i].Val == q[size-i-1].Val {
				q = append(q, q[i].Left, q[i].Right)
			} else {
				return false
			}
		}
		q = q[size:]
	}
	return true
}
