/*
Given a square array of integers A, we want the minimum sum of a falling path through A.

A falling path starts at any element in the first row, and chooses one element from each row.
The next row's choice must be in a column that is different from the previous row's column by at most one.



Example 1:

Input: [[1,2,3],[4,5,6],[7,8,9]]
Output: 12
Explanation:
The possible falling paths are:
[1,4,7], [1,4,8], [1,5,7], [1,5,8], [1,5,9]
[2,4,7], [2,4,8], [2,5,7], [2,5,8], [2,5,9], [2,6,8], [2,6,9]
[3,5,7], [3,5,8], [3,5,9], [3,6,8], [3,6,9]
The falling path with the smallest sum is [1,4,7], so the answer is 12.



Constraints:

1 <= A.length == A[0].length <= 100
-100 <= A[i][j] <= 100
*/

func Make2D(r,c, v int) [][] int {
    ret := make([][]int, r)
    for i:=0; i < len(ret); i++ {
        ret[i] = make([]int, c)
        if v > 0 {
            for j := 0; j < c; j++ {
                ret[i][j] = v
            }
        }
    }
    return ret
}

func Min(x,y int) int {
    if x > y {return y}
    return x
}

func MinPrevDP(DP []int, j int) int {
    ret := DP[j]
    if j > 0 {
        ret = Min(ret, DP[j-1])
    }
    if j < len(DP)-1 {
        ret = Min(ret, DP[j+1])
    }
    return ret
}

func minFallingPathSum(A [][]int) int {
    if (len(A) == 0) {return 0}
    DP := Make2D(len(A), len(A[0]),0)
    ret := 200
    for i, R := range(A) {
        for j, Aij := range(R) {
            if i == 0 {
                DP[i][j] = Aij
                continue
            }
            DP[i][j] = MinPrevDP(DP[i-1], j) + Aij
        }
    }
    for _,x := range(DP[len(A)-1]) {
        ret = Min(ret, x)
    }
    return ret
}
