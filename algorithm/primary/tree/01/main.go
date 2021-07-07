package _01

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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var deep int
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		deep++
		queue = queue[size:]
	}
	return deep
}
