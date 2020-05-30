/*
Path Sum 1

Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func IB(root *TreeNode, sum int, found bool) bool {
    if (found) {
        return found
    }

    if root == nil {return false}
    if root.Left == nil && root.Right == nil {
        x := sum - root.Val
        return x == 0
    }

    found = IB(root.Left, sum - root.Val, found)
    if !found {
        found = IB(root.Right, sum - root.Val, found)
    }
    return found
}

func hasPathSum(root *TreeNode, sum int) bool {
    return IB(root, sum, false)
}
/*
Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \    / \
7    2  5   1
Return:

[
   [5,4,11,2],
   [5,8,4,5]
]
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IB(root *TreeNode, sum int, tmp []int, r [][]int) [][]int {

    if root == nil {return r}
    tmp = append(tmp, root.Val)
    if root.Left == nil && root.Right == nil {
        x := sum - root.Val
        if x == 0 {
            r = append(r, append([]int(nil), tmp...))
        }
        return r
    }

    r = IB(root.Left, sum - root.Val, tmp, r)
    r = IB(root.Right, sum - root.Val, tmp, r)
    return r
}
func pathSum(root *TreeNode, sum int) (r [][]int) {
    r = IB(root, sum,[]int{}, r)
    //fmt.Println(r)
    return r
}

/*
Path Sum 3
You are given a binary tree in which each node contains an integer value.

Find the number of paths that sum to a given value.

The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

Example:

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

Return 3. The paths that sum to 8 are:

1.  5 -> 3
2.  5 -> 2 -> 1
3. -3 -> 11

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IB(root *TreeNode, sum int,r int) int {

    if root == nil {return r}
    //fmt.Println(acc == root.Val, root, sum, acc, r)
    if (sum == root.Val) {
        r++
    }
    r = IB(root.Left, sum - root.Val, r)
    r = IB(root.Right, sum - root.Val, r)
    return r
}
func pathSum(root *TreeNode, sum int) int {
    r := IB(root, sum, 0)
    if root == nil {
        return r
    }
    if root.Left != nil{
        r += pathSum(root.Left, sum)
    }
    if root.Right  != nil{
        r += pathSum(root.Right, sum)
    }
    return r
}

