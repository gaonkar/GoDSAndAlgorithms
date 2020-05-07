Check if a linked list is palindrom


/*
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

func isPalindrome(head *ListNode) bool {
    s:=head
    f:=head
    p:= head
    // find midpoint
    if (f == nil) {return true}
    if (f.Next == nil) { return true}
    for (f  != nil) {
        p = s
        s = s.Next
        f = f.Next
        if (f != nil) {
            f = f.Next
        }
    }
    // break the list and reverse midpoint
    p.Next = nil
    f = reverseList(s)
    s = head
    // check if they match
    for (s != nil && f != nil) {
        if (s.Val != f.Val) { return false}
        s = s.Next
        f = f.Next
    }
    return true
}


