/*
143. Reorder List
Medium

1663

109

Add to List

Share
Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example 1:

Given 1->2->3->4, reorder it to 1->4->2->3.
Example 2:

Given 1->2->3->4->5, reorder it to 1->5->2->4->3.

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

    if (head == nil || head.Next == nil) { return head}

    prev:= head
    cur:= prev.Next
    prev.Next = nil
    next:= cur.Next
    for next != nil {
        cur.Next = prev
        prev = cur
        cur = next
        next = cur.Next
    }
    cur.Next = prev
    return cur
}

func reorderList(head *ListNode)  {
    S := head
    s := S
    f := S
    p := S
    c := 1
    if head == nil {return}
    //find the midpoint
    for f.Next != nil {

        f = f.Next
        c++
        if (f.Next == nil) {
            break
        }
        p = s
        s = s.Next
        f = f.Next
        c++
    }
    if c % 2 == 0 {
        p = s
        s = s.Next
        p.Next = nil
    } else {

    }
    fmt.Println(s, p, c)
    s1 := reverseList(s)
    s2 := S
    // weave them alternatively
    for s1 != nil {
        n1 := s1.Next
        n2 := s2.Next
        //fmt.Println(s1,s2, n1, n2)
        s2.Next = s1
        s1.Next = n2
        s1 = n1
        s2 = n2
    }

}
