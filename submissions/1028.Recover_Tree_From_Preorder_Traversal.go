/*
1028. Recover a Tree From Preorder Traversal
We run a preorder depth first search on the root of a binary tree.
At each node in this traversal, we output D dashes (where D is the depth of this node),
then we output the value of this node.  (If the depth of a node is D, the depth of its
immediate child is D+1.  The depth of the root node is 0.)

If a node has only one child, that child is guaranteed to be the left child.

Given the output S of this traversal, recover the tree and return its root.

Input: "1-2--3---4-5--6---7"
Output: [1,2,5,3,null,6,null,4,null,7]
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func BT(arr []int, lvl []int, c int, r *TreeNode) *TreeNode {
    var RT TreeNode
    var LT TreeNode
    r.Val = arr[0]
    L :=1
    for L < len(arr) && lvl[L] != c {L++}
    H := L+1
    for H < len(arr) && lvl[H] != c {H++}
    //fmt.Println(L,H, arr, lvl)
    if H >= len(arr) && L >= len(arr) { return r}
    if H < len(arr) {
        r.Right = BT(arr[H:], lvl[H:], c+1, &RT)
    } else {
        H = len(arr)
    }
    r.Left = BT(arr[L:H], lvl[L:H], c+1, &LT)
    return r
}


func recoverFromPreorder(S string) *TreeNode {
    var arr []int
    var lvl []int
    L := 0
    i := 0
    var R TreeNode

    for i < len(S) {
        j := i
        for i < len(S) && S[i] != '-' { i++}
        num, err:= strconv.Atoi(S[j:i])
        if err == nil {
            arr = append(arr, num)
            lvl = append(lvl, L)
        }
        L = 0
        for i < len(S) && S[i] == '-' {
            i++
            L++
        }
    }
    return BT(arr, lvl, 1, &R)
}
