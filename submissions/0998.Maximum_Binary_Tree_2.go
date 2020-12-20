/*
We are given the root node of a maximum tree: a tree where every node has a value greater than any other value in its subtree.

Just as in the previous problem, the given tree was constructed from an list A (root = Construct(A)) recursively with the following Construct(A) routine:

If A is empty, return null.
Otherwise, let A[i] be the largest element of A.  Create a root node with value A[i].
The left child of root will be Construct([A[0], A[1], ..., A[i-1]])
The right child of root will be Construct([A[i+1], A[i+2], ..., A[A.length - 1]])
Return root.
Note that we were not given A directly, only a root node root = Construct(A).

Suppose B is a copy of A with the value val appended to it.  It is guaranteed that B has unique values.

Problem 623 construct binary tree.
So flatten the tree, append val and construct

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    stk := make([]*TreeNode, 0)
    for _,x := range(nums) {
        n := &TreeNode{x, nil, nil}
        for len(stk) > 0 && stk[len(stk)-1].Val < x {
            n.Left = stk[len(stk)-1]
            stk = stk[:len(stk)-1]
        }
        if len(stk) > 0 {
            stk[len(stk)-1].Right = n
        }
        stk = append(stk, n)
    }
    return stk[0]
}
func flattenTree(root *TreeNode) []int {
    if root == nil {return []int{}}
    arr := flattenTree(root.Left)
    arr = append(arr, root.Val)
    arr = append(arr, flattenTree(root.Right)...)
    return arr
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
    return constructMaximumBinaryTree(append(flattenTree(root), val))
}
