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

// 前序遍历
func TreeToPreorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var nums []int
	nums[0] = root.Val

	// 只有一个节点
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

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

	res = TreeToInorder(root.Left)
	res = append(res, root.Val)
	res = append(res, TreeToInorder(root.Right)...)
	return res 
}
