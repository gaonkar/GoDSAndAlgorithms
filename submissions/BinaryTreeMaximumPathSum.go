/*
Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to
any node in the tree along the parent-child connections. The path must contain at least
one node and does not need to go through the root.

Example 1:

Input: [1,2,3]

       1
      / \
     2   3

Output: 6
Example 2:

Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

Output: 42
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max (x int, y int) int {
    if (x < y) {return y}
    return x
}

var Mtop int

func MPath(root *TreeNode) int {
    if (root == nil) {
        return 0
    }
    L:= MPath(root.Left)
    R:= MPath(root.Right)
    // This case it is either the Left or Right subtree with the root or the root itself
    MS:= max(max(L,R)+root.Val, root.Val)
    // include a possibility that root is part of the Max Sum
    MT := max(MS, L + R + root.Val)
    // store the best we have found
    Mtop = max(Mtop, MT)
    return MS
}

func maxPathSum(root *TreeNode) int {
    Mtop = -100000
    MPath(root)
    return Mtop
}
