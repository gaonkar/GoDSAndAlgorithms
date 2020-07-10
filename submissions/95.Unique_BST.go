/*
Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.

Example:

Input: 3
Output:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
Explanation:
The above output corresponds to the 5 unique BST's shown below:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

*/
func MakeNode(Val int, l, r *TreeNode) *TreeNode {
    var R TreeNode
    R.Val = Val
    R.Left = l
    R.Right = r
    return &R
}


func GenBST(N []int) []*TreeNode {
    arr := make([]*TreeNode, 0)
    if len(N) == 0 {return nil}
    if len(N) == 1 {
        arr = append(arr, MakeNode(N[0],nil,nil))
        return arr
    }
    //N[0]
    r := GenBST(N[1:])
    for j:= range(r) {
        arr = append(arr, MakeNode(N[0],nil,r[j]))
    }
    if len(N) > 0 {
        l := GenBST(N[:len(N)-1])
        for i:= range(l) {
            arr = append(arr, MakeNode(N[len(N)-1],l[i],nil))
        }
    }
    for x:=1; x < len(N)-1;x++ {
        fmt.Println(N[x], N[:x], N[x+1:])
        l := GenBST(N[:x])
        r := GenBST(N[x+1:])
        //fmt.Println(l,r)
        for i:= range(l) {
            for j:= range(r) {
                arr = append(arr, MakeNode(N[x],l[i],r[j]))
            }
        }
    }
    return arr
}

func generateTrees(n int) []*TreeNode {
    N := make([]int, n)
    for i:=0; i < n; i++ {
        N[i] = i+1
    }
    return GenBST(N)
}
