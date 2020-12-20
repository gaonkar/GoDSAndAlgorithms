/*
Given an array of unique integers, each integer is strictly greater than 1.

We make a binary tree using these integers and each number may be used for any number of times.

Each non-leaf node's value should be equal to the product of the values of it's children.

How many binary trees can we make?  Return the answer modulo 10 ** 9 + 7.

Example 1:

Input: A = [2, 4]
Output: 3
Explanation: We can make these trees: [2], [4], [4, 2, 2]
Example 2:

Input: A = [2, 4, 5, 10]
Output: 7
Explanation: We can make these trees: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].


Note:

1 <= A.length <= 1000.
2 <= A[i] <= 10 ^ 9.
*/
func numFactoredBinaryTrees(A []int) int {
    sort.Ints(A)
    dp := make([]int, len(A))
    IDX := make(map[int] int)
    for i:=0; i < len(A); i++ {
        dp[i] = 1
        IDX[A[i]]=i+1 // adding one so that we can find it use > 0
    }
    MOD := 1000000007
    for i:=0; i < len(A); i++ {
        for j:=0; j < i; j++ {
            if A[i] % A[j] != 0 {continue}
            // A[j] is the left child, find the right
            right := A[i]/ A[j]
            idx := IDX[right]
            if idx > 0 {
                idx--
                dp[i] = (dp[i] + dp[j]*dp[idx]) % MOD
            }
        }
    }
    ret := 0
    for _,x := range(dp) {
        ret = (ret + x)%MOD
    }
    return ret
}
