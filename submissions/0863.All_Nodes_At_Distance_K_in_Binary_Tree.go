/*
We are given a binary tree (with root node root), a target node, and an integer value K.

Return a list of the values of all nodes that have a distance K from the target node.  The answer can be returned in any order.



Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2

Output: [7,4,1]

Explanation:
The nodes that are a distance 2 from the target node (with value 5)
have values 7, 4, and 1.



Note that the inputs "root" and "target" are actually TreeNodes.
The descriptions of the inputs above are just serializations of these objects.


Note:

The given tree is non-empty.
Each node in the tree has unique values 0 <= node.val <= 500.
The target node is a node in the tree.
0 <= K <= 1000.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// find all children at distance k
func distanceAsRoot(root *TreeNode, K int) (r []int) {
    var helper func(root *TreeNode, K int)

    helper = func(root *TreeNode, K int) {
        if root == nil {return}
        if K == 0 {
            r = append(r, root.Val)
            return
        }
        helper(root.Left, K-1)
        helper(root.Right, K-1)
    }
    helper(root, K)
    return r
}

type Path struct {
    r *TreeNode
    isleft bool
}

func PathToTarget(root *TreeNode, target *TreeNode) (r []*Path)  {
    var helper func(root *TreeNode, target *TreeNode) bool

    helper = func(root *TreeNode, target *TreeNode) bool {
        if root == nil {return false}
        if root.Val == target.Val {
            return true
        }
        path := Path{root, true}
        r = append(r, &path)
        found := helper(root.Left, target)
        if found  {return found}
        path.isleft = false
        found = helper(root.Right, target)
        if !found {
            r = r[:len(r)-1]
        }
        return found
    }
    found := helper(root, target)
    if found {
        return r
    }
    return nil
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
    path := PathToTarget(root, target)
    r := distanceAsRoot(target, K)
    L := len(path)-1
    K-- // the path has list of parents
    for K >= 0 &&  L >= 0 {
        p := path[L]

        if K == 0 {
            r = append(r, p.r.Val)
        } else {
            if p.isleft {
                target = p.r.Right
            } else {
                target = p.r.Left
            }
            r = append(r, distanceAsRoot(target, K-1)...)
        }
        K--
        L--
    }
    return r
}
