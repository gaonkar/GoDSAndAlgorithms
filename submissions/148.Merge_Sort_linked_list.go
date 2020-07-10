
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head:= &ListNode{0, nil}
    if (l1 == nil) {return l2}
    if (l2 == nil) { return l1}
    mx := head
    for (l1 != nil && l2 != nil) {
        if (l1.Val < l2.Val) {
            mx.Next = l1
            l1 = l1.Next
        } else {
            mx.Next = l2
            l2 = l2.Next
        }
        mx = mx.Next
    }
    if (l1 == nil) {
        mx.Next = l2
    } else {
        mx.Next = l1
    }
    return head.Next
}

func Count(head *ListNode) (c int) {
    for head != nil {
        head = head.Next
        c++
    }
    return c
}

func Split(l *ListNode, n int) *ListNode {
    p := l
    l = l.Next
    for n > 1 {
        p = l
        l = l.Next
        n--
    }
    p.Next = nil
    return l
}

func MergeSort(l1 *ListNode, n int) *ListNode {
    if n <= 1 {
        return l1
    }
    m := n/2
    l2 := Split(l1, m)
    l2 = MergeSort(l2, n-m)
    l1 = MergeSort(l1, m)
    return mergeTwoLists(l1,l2)
}


func sortList(head *ListNode) *ListNode {
    N := Count(head)
    fmt.Println(N)
    return MergeSort(head, N)
}
