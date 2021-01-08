/*
Given a binary tree where node values are digits from 1 to 9. A path in the binary tree is said to be pseudo-palindromic
if at least one permutation of the node values in the path is a palindrome.
Return the number of pseudo-palindromic paths going from the root node to leaf nodes.

Example 1:

Input: root = [2,3,1,3,1,null,1]
Output: 2
Explanation: There are three paths going from the root node to leaf nodes: the red path [2,3,3], the green path [2,1,1], 
and the path [2,3,1]. Among these paths only red path and green path are pseudo-palindromic paths since the red path 
[2,3,3] can be rearranged in [3,2,3] (palindrome) and the green path [2,1,1] can be rearranged in [1,2,1] (palindrome).

Example 2:

Input: root = [2,1,1,1,3,null,null,null,null,null,1]
Output: 1
Explanation: There are three paths going from the root node to leaf nodes: the green path [2,1,1], the path [2,1,3,1], 
and the path [2,1]. Among these paths only the green path is pseudo-palindromic since [2,1,1] can be rearranged in 
[1,2,1] (palindrome).
Example 3:

Input: root = [9]
Output: 1


Constraints:

The given binary tree will have between 1 and 10^5 nodes.
Node values are digits from 1 to 9.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsPalindrome(s []byte) bool {
    B := make([]int, 10)
    L := len(s)
    for i := 0; i < L; i++ {
        B[s[i]]++
    }
    odd := 0
    //fmt.Println(B, s)
    for i := 0; i < len(B); i++ {
        if B[i] % 2 != 0 {
            odd++
            //string is even
            if L % 2 == 0 {return false}
        }
        if odd > 1 {return false}
    }
    return true
}
func IOT(root *TreeNode, s []byte, l int) int{
    if root == nil {
        return 0
    }
    s = append(s, byte(root.Val))
    ret := 0
    //fmt.Println(root.Val, l, s)
    if root.Left == nil && root.Right == nil {
        //fmt.Println("IP",root.Val)
        if (IsPalindrome(s[1:])) {
            ret++
        }
        s = s[:len(s)-1]
        return ret
    }
    if root.Left != nil {
        ret += IOT(root.Left, s, l + 1)
    }
    if root.Right != nil{
        ret += IOT(root.Right, s, l + 1)
    }
    s = s[:len(s)-1]
    return ret
}
func pseudoPalindromicPaths (root *TreeNode) int {
    s := make([]byte, 1)
    return IOT(root, s, 0)
}
