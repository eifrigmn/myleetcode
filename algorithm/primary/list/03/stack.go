package _03

type stack struct {
	arr []*ListNode
}

func reverseList_stack(head *ListNode) *ListNode {
	s := &stack{
		arr: []*ListNode{},
	}
	for head != nil {
		tmp := head.Next
		head.Next = nil
		s.push(head)
		head = tmp
	}

	var newHead = new(ListNode)
	h := newHead
	node := s.pop()
	for node != nil {
		h.Next = node
		h = h.Next
		node = s.pop()
	}
	return newHead.Next
}

func (s *stack) push(node *ListNode) {
	s.arr = append(s.arr, node)
}

func (s *stack) pop() *ListNode {
	if len(s.arr) == 0 {
		return nil
	}
	node := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return node
}
