package linkedList

type Item interface{}

// Node 结点信息
type Node struct {
	data Item
	next *Node
}
