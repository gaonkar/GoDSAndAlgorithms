/*

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

*/

func isBST(root *TreeNode, min map[*TreeNode] int, max map[*TreeNode] int) bool {
    if (root == nil) {return true}
    L := (root.Left == nil)
    R := (root.Right == nil)
    if !L {
        L = (root.Val > max[root.Left])
    }
    if !R {
        R = (root.Val < min[root.Right])
    }
    if L {
        L = isValidBST(root.Left)
    }
    if R {
        R = isValidBST(root.Right)
    }
    return L && R
}

func Min(x int, y int) int {
    if (x < y) {return x}
    return y
}

func Max(x int, y int) int {
    if (x > y) {return x}
    return y
}

func MakeMinMap(root *TreeNode, min map[*TreeNode] int, max map[*TreeNode] int) {
    if root == nil {return}
    if root.Left == nil && root.Right == nil{
        max[root] = root.Val
        min[root] = root.Val
        return
    }
    if root.Right == nil {
        max[root] = root.Val
        MakeMinMap(root.Left, min, max)
        min[root] = Min(root.Val, min[root.Left])
        return
    }
    if root.Left == nil {
        min[root] = root.Val
        MakeMinMap(root.Right, min, max)
        max[root] = Max(root.Val, max[root.Right] )
        return
    }
    MakeMinMap(root.Left, min, max)
    min[root] = Min(root.Val, min[root.Left])
    MakeMinMap(root.Right, min, max)
    max[root] = Max(root.Val, max[root.Right])
}

func isValidBST(root *TreeNode) bool {
    min := make(map[*TreeNode] int)
    max := make(map[*TreeNode] int)
    MakeMinMap(root, min, max)
    return isBST(root, min, max)
}

