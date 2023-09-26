package linkedlist

type ListNode struct {
	Next *ListNode
	Val  int
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}

	dummy.Next = head
	prev := dummy
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			v := cur.Val
			for cur != nil && cur.Val == v {
				cur = cur.Next
			}
			prev.Next = cur
		} else {
			prev = cur
			cur = cur.Next
		}
	}

	return dummy.Next
}
