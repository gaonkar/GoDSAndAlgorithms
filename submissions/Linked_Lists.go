/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * Delete all duplicate values and only leave distinct values
 */
func deleteDuplicates(head *ListNode) *ListNode {

	h := head
	if h == nil || h.Next == nil {
		return head
	}
	n := h.Next
	if h.Val == n.Val {
		// the head has to be removed
		for n != nil && h.Val == n.Val {
			n = n.Next
		}
		h = deleteDuplicates(n)
	} else {
		h.Next = deleteDuplicates(n)
	}
	return h
}

/*
Flatten a Multilevel Doubly Linked List


You are given a doubly linked list which in addition to the next and previous pointers,
it could have a child pointer, which may or may not point to a separate doubly linked list.
These child lists may have one or more children of their own, and so on, to produce a
multilevel data structure, as shown in the example below.

Flatten the list so that all the nodes appear in a single-level, doubly linked list.
You are given the head of the first level of the list.
*/

func FindLast(root *Node) *Node {
	for root != nil && root.Next != nil {
		root = root.Next
	}
	return root
}

func Flatten(root *Node) (r *Node) {
	r = root
	for root != nil {
		if root.Child != nil {
			child := root.Child
			child.Prev = root
			next := root.Next
			last := FindLast(child)
			last.Next = next
			//fmt.Println(root, child, last)
			if next != nil {
				next.Prev = last
			}
			root.Next = child
			root.Child = nil
		}
		root = root.Next
	}
	root = r
	for root != nil && root.Next != nil {
		//fmt.Println(root)
		root = root.Next
	}
	return r
}

func flatten(root *Node) *Node {
	return Flatten(root)
}
