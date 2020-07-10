/*
Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
*/
func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	prev := head
	re, rs := head, head
	ht := head
	for m > 1 {
		ht = prev
		prev = prev.Next
		rs = prev
		m--
		n--
	}
	//fmt.Println(head,rs,prev)
	cur := prev.Next
	rs.Next = nil
	next := cur.Next
	for n > 1 {
		//fmt.Println("*", prev,cur,next,n)
		cur.Next = prev
		prev = cur
		re = cur
		cur = next
		if cur == nil {
			break
		}
		next = cur.Next
		n--

	}
	ht.Next = re
	rs.Next = cur
	return head
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}
	var Dummy ListNode
	Dummy.Next = head
	// dummy node helps get rid of corner cases
	ReverseBetween(&Dummy, m+1, n+1)
	return Dummy.Next
}
