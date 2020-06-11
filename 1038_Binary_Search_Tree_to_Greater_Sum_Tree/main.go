package _1038

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

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	res := inorder(root)
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			res[i].Val = res[i].Val + res[j].Val
		}
	}
	return root
}

func inorder(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []*TreeNode{root}
	}
	var res []*TreeNode

	res = inorder(root.Left)
	res = append(res, root)
	res = append(res, inorder(root.Right)...)
	return res
}
