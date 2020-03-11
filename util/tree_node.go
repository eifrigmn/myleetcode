package util

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63

func IntsToTreeNode(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	for i := 1; i < n; {
		node := queue[0]
		queue = queue[1:]

		if i < n && nums[i] != NULL {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && nums[i] != NULL {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// Tree2ints 把 *TreeNode 按照行还原成 []int
func TreeToInts(tn *TreeNode) []int {
	res := make([]int, 0, 1024)

	queue := []*TreeNode{tn}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			nd := queue[i]
			if nd == nil {
				res = append(res, NULL)
			} else {
				res = append(res, nd.Val)
				queue = append(queue, nd.Left, nd.Right)
			}
		}
		queue = queue[size:]
	}

	i := len(res)
	for i > 0 && res[i-1] == NULL {
		i--
	}

	return res[:i]
}

// 前序遍历
func TreeToPreorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// 只有一个节点
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	nums := []int{root.Val}
	nums = append(nums, TreeToPreorder(root.Left)...)
	nums = append(nums, TreeToPreorder(root.Right)...)
	return nums
}

// 中序遍历
func TreeToInorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var res []int
	res = TreeToInorder(root.Left)
	res = append(res, root.Val)
	res = append(res, TreeToInorder(root.Right)...)
	return res
}

func TreeToPostorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := TreeToPostorder(root.Left)
	res = append(res, TreeToPostorder(root.Right)...)
	res = append(res, root.Val)
	return res
}
