/*
105. Construct Binary Tree from Preorder and Inorder Traversal
Medium

2989

86

Add to List

Share
Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
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
func buildTree(P []int, I []int) *TreeNode {
    if len(P) == 0 {return nil}
    r :=TreeNode{P[0], nil, nil}
    i := 0
    if len(P) == 1 {
        // assert P[0] = I[0]
        return &r
    }
    for i < len(I) && r.Val != I[i]{
        i++
    }
    //fmt.Println(P, I, i)

    PL := P[1:i+1]
    IL := I[0:i]
    PR := P[i+1:]
    IR := I[i+1:]
    //fmt.Println(PL, PR, IL, IR)
    if len(IL) > 0 {
        r.Left = buildTree(PL, IL)
    }
    if len(IR) > 0 {
        r.Right = buildTree(PR, IR)
    }
    return &r
}
