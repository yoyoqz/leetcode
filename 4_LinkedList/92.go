package linkedlist

type ListNode struct {
	Next *ListNode
	Val  int
}

/*
func reverse(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur

		cur = next
	}

	return prev
}

*/

func reverseBetween(head *ListNode, m int, n int) *ListNode {

}
