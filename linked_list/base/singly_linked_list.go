package linked_list

// SSL singly linked list
type SSL struct {
	head *Node
	len  int
}

// Back 尾结点
func (l *SSL) Back() *Node {
	curr := l.head
	for curr != nil && curr.next != nil {
		curr = curr.next
	}
	return curr
}

// Front 头结点
func (l *SSL) Front() *Node {
	return l.head
}

// Append 后向添加结点
func (l *SSL) Append(n *Node) {
	// 空链表
	if l.head == nil {
		l.head = n
	} else {
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = n
	}
	l.len++
}

// Prepend 前向添加结点
func (l *SSL) Prepend(n *Node) {
	if l.head == nil {
		l.head = n
	} else {
		n.next = l.head
		l.head = n
	}
	l.len++
}

// InsertBefore 在指定结点dest前插入结点n
// 需要找到指定结点的上一个结点，然后进行插入
func (l *SSL) InsertBefore(n, dest *Node) {
	if l.head == dest {
		// 指定结点为头结点
		n.next = dest
		l.head = n
		l.len++
	} else {
		// 遍历寻找指定结点的前序结点
		curr := l.head
		for curr != nil && curr.next != nil && curr.next != dest {
			curr = curr.next
		}
		// 指定结点在链表中，才插入，否则不做操作
		if curr.next == dest {
			n.next = dest
			curr.next = n
			l.len++
		}
	}
}

// InsertAfter 在指定结点dest后插入结点n
func (l *SSL) InsertAfter(n, dest *Node) {
	curr := l.head
	// 遍历链表寻找指定结点dest并将其赋值给curr,如果未找到之地结点，则将尾结点赋值给curr
	for curr != dest && curr != nil && curr.next != nil {
		curr = curr.next
	}
	if curr == dest {
		n.next = dest.next
		curr.next = n
		l.len++
	}
}

// Remove 移除指定结点
func (l *SSL) Remove(n *Node) {
	// 指定结点为头结点
	if l.head == n {
		l.head = n.next
		n.next = nil
		l.len--
	} else {
		// 搜索指定结点的前序结点
		curr := l.head
		for curr != nil && curr.next != nil && curr.next != n {
			curr = curr.next
		}
		if curr.next == n {
			curr.next = n.next
			n.next = nil
			l.len--
		}
	}
}

// RemoveBefore 移除指定结点前的那个结点
func (l *SSL) RemoveBefore(n *Node) {
	if l.head == nil || l.head == n {
		return
	}

	if l.head.next == n {
		l.head = n
		l.len--
	} else {
		curr := l.head
		for curr.next.next != n && curr.next.next != nil {
			curr = curr.next
		}

		if curr.next.next == n {
			curr.next = n
			l.len--
		}
	}

}

// RemoveAfter 移除指定结点后的那个结点
func (l *SSL) RemoveAfter(n *Node) {
	if l.head == nil {
		return
	}

	curr := l.head
	for curr != n && curr != nil && curr.next != nil {
		curr = curr.next
	}

	if curr == n && curr.next != nil {
		curr.next = n.next
	}
}

// GetAtPos 获取链表中索引为i的结点
func (l *SSL) GetAtPos(idx int) *Node {
	curr := l.head
	for i :=0; i<idx;i++{
		if curr.next == nil {
			return nil
		}
		curr = curr.next
	}
	return curr
}

// Find 获取链表中指定数据值的结点信息
func (l *SSL) Find(val interface{}) *Node {
	curr := l.head
	for curr != nil && curr.next != nil && curr.data != val {
		curr = curr.next
	}
	if curr.data == val {
		return curr
	}
	return nil
}

// Length 链表中的结点数量
func (l *SSL) Length() int {
	return l.len
}
