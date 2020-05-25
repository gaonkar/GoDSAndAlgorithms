

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func FindSmallest(lists []*ListNode) *ListNode {

    sNode := lists[0]
    i := 0
    // find the non-null entry
    for i< len(lists) {
        sNode = lists[i]
        if sNode != nil {
            break
        }
        i++
    }
    if (sNode == nil) {return nil}
    j := i
    i++

    for i < len(lists) {
        if lists[i] != nil && sNode.Val > lists[i].Val {
            sNode = lists[i]
            j = i
        }
        i++
    }
    //fmt.Println(sNode, j)
    if (lists[j] != nil) {
        lists[j] = lists[j].Next
    }
    return sNode
}

func mergeKLists(lists []*ListNode) *ListNode {
    head := FindSmallest(lists)
    r := head
    for r != nil {
        r.Next = FindSmallest(lists)
        r = r.Next
    }
    return head
}

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

func MKL(lists[] *ListNode, l int, h int) *ListNode {
    if l == h {return lists[h]}
    if l+1 == h {return mergeTwoLists(lists[l],lists[h])}
    m := (h + l )/2
    L := MKL(lists, l, m)
    R := lists[m+1]
    if (m+1 < h) {
        R = MKL(lists,m+1,h)
    }
    return mergeTwoLists(L,R)

}

func mergeKLists(lists []*ListNode) *ListNode {
    L := len(lists)
    if L == 0 {return nil}
    return MKL(lists, 0, len(lists)-1)
}
