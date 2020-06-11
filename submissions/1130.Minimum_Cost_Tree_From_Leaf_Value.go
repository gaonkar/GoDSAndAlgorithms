/*
Given an array arr of positive integers, consider all binary trees such that:

Each node has either 0 or 2 children;
The values of arr correspond to the values of each leaf in an in-order traversal of the tree.  (Recall that a node is a leaf if and only if it has 0 children.)
The value of each non-leaf node is equal to the product of the largest leaf value in its left and right subtree respectively.
Among all possible binary trees considered, return the smallest possible sum of the values of each non-leaf node.  It is guaranteed this sum fits into a 32-bit integer.



Example 1:

Input: arr = [6,2,4]
Output: 32
Explanation:
There are two possible trees.  The first has non-leaf node sum 36, and the second has non-leaf node sum 32.

    24            24
   /  \          /  \
  12   4        6    8
 /  \               / \
6    2             2   4


Constraints:

2 <= arr.length <= 40
1 <= arr[i] <= 15
It is guaranteed that the answer fits into a 32-bit signed integer (ie. it is less than 2^31).

*/

func Min(x, y int) int {
    if x < 0 {return y}
    if x < y {return x}
    return y
}
func Max(x, y int) int {
    if x < y {return y}
    return x
}

/*
 Memoization
 */
func MinT(A []int, D, M [][]int, l, h int) int{
    if h == l {return 0}
    if D[l][h] >= 0 {return D[l][h]}
    m := -1

    for i:=l; i < h; i++ {
        P := MinT(A,D,M, l, i)
        A := MinT(A, D,M, i+1, h)
        x :=  P + A + M[l][i] * M[i+1][h]
        //fmt.Println(l,i,h, M[l][i], M[i+1][h], P, A)
        m = Min(m, x)
    }
    D[l][h] = m
    return m
}


func mctFromLeafValues(A []int) int {
    L:= len(A)
    M:=make([][]int, L)
    D:=make([][]int, L+1)
    for i:= range(D) {

        D[i] = make([]int, L+1)

        for j:=range(D[i]) {
            D[i][j] = -1
        }
        if i == L {continue}
        M[i] = make([]int, L)
        m := A[i]
        for j := i; j < L; j++ {
            m = Max(A[j], m)
            M[i][j] = m
        }
    }

    y := MinT(A, D,M, 0, L-1)
    //for i:=0; i < L+1;i++ {    fmt.Println(D[i])}
    return y
}
