/*
Given inorder and postorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

inorder = [9,3,15,20,7]
postorder = [9,15,7,20,3]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Index(a []int, val int) int {
    i := 0
    for i < len(a) && val != a[i]{i++}
    return i
}

/*
 In inorder vs post order, the inorder index decides the partition of the tree
 */

func buildTree(inorder []int, postorder []int) *TreeNode {
    var R TreeNode
    Pl := len(postorder)
    if Pl == 0 {return nil}
    if Pl == 1 {
        R.Val = postorder[0]
        return &R
    }

    last := postorder[len(postorder)-1]
    R.Val = last
    i := Index(inorder, last)
    R.Left = buildTree(inorder[:i], postorder[:i])
    R.Right = buildTree(inorder[i+1:], postorder[i:Pl-1])
    return &R
}
