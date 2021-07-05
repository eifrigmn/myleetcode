package _04

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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var q = []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		var m []int
		for i := 0; i < size; i++ {
			if q[i] != nil {
				m = append(m, q[i].Val)
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		result = append(result, m)
		q = q[size:]
	}
	return result
}
