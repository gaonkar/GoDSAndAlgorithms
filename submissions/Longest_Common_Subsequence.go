/*
Given two strings text1 and text2, return the length of their longest common subsequence.

A subsequence of a string is a new string generated from the original string with some characters(can be none)
deleted without changing the relative order of the remaining characters. (eg, "ace" is a subsequence of "abcde"
while "aec" is not). A common subsequence of two strings is a subsequence that is common to both strings.

Constraints:

1 <= text1.length <= 1000
1 <= text2.length <= 1000
The input strings consist of lowercase English characters only.

Solution for the sequence itself
Dynamic Programming
LCS[X(i),X(j)]  = 0, if i = 0 or j = 0
LCS[X(i), X(j)] = CONCAT{LCS[X(i-1), X(j-1)],X(i)} if X(i) = X(j) i>0,j>0
LCS[X(i), X(j)] = MAX{LCS[X(i), X(j-1)],LCS[X(i-1), X(j)]} if X(i) != X(j) i>0,j>0

Solution for the length
Dynamic Programming
LCS[X(i),X(j)]  = 0, if i = 0 or j = 0
LCS[X(i), X(j)] = LCS[X(i-1), X(j-1)] + 1 if X(i) = X(j) i>0,j>0
LCS[X(i), X(j)] = MAX{LCS[X(i), X(j-1)],LCS[X(i-1), X(j)]} if X(i) != X(j) i>0,j>0



*/
import "fmt"
func Max(x, y int) int{
    if x < y {
        return y
    }
    return x
}

func longestCommonSubsequence(text1 string, text2 string) int {
    var curr = make([]int, len(text1)+1)
    var prev = make([]int, len(text1)+1)
    //fmt.Println(curr, prev)
    for j := range text2 {
        for i:= range text1{
            if text1[i] == text2[j] {
                curr[i+1] = prev[i] + 1
            } else {
                curr[i+1] = Max(curr[i], prev[i+1])
            }
            //fmt.Println(i,j, curr)
        }
        prev = curr
        curr = make([]int, len(text1)+1)
    }
    return prev[len(text1)]
}

/*
 Apparently, the for loop can be replaced with bin search. But really non-intuitive. Note how J is initialized to M each time
*/
func lengthOfLIS(nums []int) int {
    L := len(nums)
    if L == 0 { return 0}
    P := make([]int, L)
    S := make([]int, L+1)
    LL:= make([]int, L)
    j := 1
    S[0] = 0
    LL[0]=1
    M := 0
    for i := 1; i < L ; i++ {
        j = M
        for j >= 0 && nums[S[j]] >= nums[i] {j--}
        j++
        if j == 0 {
            P[i]=i
            LL[i] = 1
        } else {
            P[i] = S[j-1]
            LL[i] = 1 + LL[P[i]]
        }
        S[j] = i
        if M < j { M = j}
        //fmt.Println(i,j, S, P, M)
    }
    fmt.Println(nums)
    fmt.Println(S)
    fmt.Println(P)
    fmt.Println(LL)
    return M+1
}
