/*
A full binary tree is a binary tree where each node has exactly 0 or 2 children.
Return a list of all possible full binary trees with N nodes.  Each element of the answer is the root node of one possible tree.
Each node of each tree in the answer must have node.val = 0.
You may return the final list of trees in any order.

Example 1:

Input: 7
Output: [[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 * Solution
 *  If N is even, its a nil set
 *  O(2^N) its a combinatorial problem
 */
func allPossibleFBT(N int) []*TreeNode {
    M := make(map[int] []*TreeNode)
    var helper func(N int) []*TreeNode
    M[0] = nil
    M[1] = make([]*TreeNode,1)
    M[1][0] = &TreeNode{0,nil,nil}

    helper = func(N int) []*TreeNode {
        if N%2 == 0 {return M[0]} // Even return nil
        v, ok := M[N]
        if ok {return v}
        ret := make([]*TreeNode, 0)
        for l := 0; l < N; l++ {
            r  := N - 1 - l     // NCK
            left := helper(l)
            right := helper(r)
            for _,x := range(left) {
                for _,y := range(right) {
                    root := &TreeNode{0,nil,nil}
                    root.Left = x
                    root.Right = y
                    ret = append(ret, root)
                }
            }
        }
        M[N] = ret
        return M[N]
    }
    helper(N)
    return M[N]
}
