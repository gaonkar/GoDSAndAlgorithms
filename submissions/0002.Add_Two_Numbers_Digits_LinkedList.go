/*

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Similar Questions**:
* [Multiply Strings (Medium)](https://leetcode.com/problems/multiply-strings/)
* [Add Binary (Easy)](https://leetcode.com/problems/add-binary/)
* [Sum of Two Integers (Easy)](https://leetcode.com/problems/sum-of-two-integers/)
* [Add Strings (Easy)](https://leetcode.com/problems/add-strings/)
* [Add Two Numbers II (Medium)](https://leetcode.com/problems/add-two-numbers-ii/)
* [Add to Array-Form of Integer (Easy)](https://leetcode.com/problems/add-to-array-form-of-integer/)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTail(t *ListNode, val int) int {
	c := 0
	s := val
	if val > 9 {
		s = val % 10
		c = 1
	}
	t.Next = &ListNode{s, nil}
	return c
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail, rem *ListNode
	var sum, carry int = 0, 0
	head = &ListNode{0, nil}
	tail = head
	for (l1 != nil) && (l2 != nil) {
		sum = l1.Val + l2.Val + carry
		carry = addTail(tail, sum)
		tail = tail.Next

		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil {
		rem = l2
	} else {
		rem = l1
	}
	for rem != nil {
		sum = rem.Val + carry
		carry = addTail(tail, sum)
		tail = tail.Next

		rem = rem.Next
	}
	if carry > 0 {
		tail.Next = &ListNode{carry, nil}
	}
	return head.Next
}
