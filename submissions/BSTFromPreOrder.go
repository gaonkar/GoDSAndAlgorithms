/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 *
 *  Return the root node of a binary search tree that matches the given preorder traversal.

 *

 */

func BSTC(V int, a []int, low int, high int) *TreeNode {
    T:= TreeNode{V, nil, nil}
    mid:= low
    for ;mid <= high; mid++ {
        if V < a[mid] {
            break
        }
    }
    if (low > high) {
        return &T
    }
    //fmt.Println(V, low, mid, high)
    if (low + 1 <= mid) {
        T.Left = BSTC(a[low], a, low + 1, mid -1)
    }
    if (mid  <= high) {
        T.Right = BSTC(a[mid], a, mid+1, high)
    }
    return &T
}

func bstFromPreorder(preorder []int) *TreeNode {
    return BSTC(preorder[0], preorder, 1, len(preorder)-1)
}
