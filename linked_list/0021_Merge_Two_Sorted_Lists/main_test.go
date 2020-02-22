package linked_list

import "testing"

func TestMergeTwoLists(t *testing.T) {
	var l1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	var l2 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	ml := mergeTwoLists(l1, l2)
	for {
		t.Log(ml.Val)
		if ml.Next == nil {
			break
		}
		ml = ml.Next
	}
}
