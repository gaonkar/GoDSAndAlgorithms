/*
Given the head of a linked list, we repeatedly delete consecutive sequences of nodes that sum to 0 until there are no such sequences.
After doing so, return the head of the final linked list.  You may return any such answer.



(Note that in the examples below, all sequences are serializations of ListNode objects.)

Example 1:

Input: head = [1,2,-3,3,1]
Output: [3,1]
Note: The answer [1,2,1] would also be accepted.
Example 2:

Input: head = [1,2,3,-3,4]
Output: [1,2,4]
Example 3:

Input: head = [1,2,3,-3,-2]
Output: [1]


Constraints:

The given linked list will contain between 1 and 1000 nodes.
Each node in the linked list has -1000 <= node.val <= 1000
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type PSum struct {
    v int
    prev *ListNode
}
func Erase(prev, curr *ListNode, val int, M map[int] PSum) *ListNode {
    ret := curr.Next
    fix := prev.Next
    prev.Next = ret

    for fix != curr {
        val += fix.Val
        delete(M, val)
        fix = fix.Next
    }
    return prev
}
func removeZeroSumSublists(head *ListNode) *ListNode {
    if head == nil {return nil}
    dumm := &ListNode{0, head}
    prev := dumm
    M := map[int] PSum{}
    sum := 0
    for head != nil {
        if head.Val == 0 {
            prev.Next = head.Next
            head = head.Next
            continue
        }
        sum += head.Val
        if sum == 0 {
            // we just recursively call from the next position as we can ignore everything till here
            return  removeZeroSumSublists(head.Next)
        }
        entry, ok := M[sum]
        if ok {
            // found a sub sequence that is adding upto 0, lets erase it
            head = Erase(entry.prev,head, sum, M)
        } else {
            M[sum] = PSum{sum, head}
        }
        prev = head
        head = head.Next

    }
    return dumm.Next
}
