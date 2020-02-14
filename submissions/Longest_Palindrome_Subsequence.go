/*
Given a string s, find the longest palindromic subsequence's length in s. You may assume that the maximum length of s is 1000. N = len(s)

Solution:
1. Its an equivalent problem of Longest Common Subsequence for s, Reverse(s)
  Time = O(N^2), Space (2N)

2. First P[i,l] represents positions in the string i to length l
0 <= i < N
1 <= l <= N
P[i,1] = 1
P[i,2] = 0 if  S[i] != S[i+1]
P[i,2] = 2 if  S[i] == S[i+1]
P[i,l] = 2 + P[i+1, l - 2]  if S[i] == S[i+l-1] 	// note that l-2, because you now look up P of len -2
P[i,l] = MAX(P[i+1,l-1], P[i, l-1]) if S[i] != S[i+l-1]

*/


import "fmt"
func Max(x, y int) int{
    if x < y {
        return y
    }
    return x
}

func Reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return
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

func LPSDynamicProgram(s string) int {
    N := len(s)
    var curr = make([]int, N)
    var prev1 = make([]int, N)
    var prev2 = make([]int, N)
    for i:= range prev1 {
        prev1[i] = 1
    }
    for l:=2; l <= N; l++ { // length of the substring
        for i:= 0; i <= N-l;i++ {
            p:=i+l-1
            if s[i] == s[p] {
                if (l == 2) {
                    curr[i] = 2
                } else {
                    curr[i] = 2 + prev2[i+1]
                }
            } else {
                curr[i] = Max(prev1[i+1], prev1[i])
            }
        }
        prev2 = prev1
        prev1 = curr
        curr = make([]int, N)
        //fmt.Println(prev1)
    }
    return prev1[0]
}

func longestPalindromeSubseq(s string) int {
    //return longestCommonSubsequence(s,Reverse(s))
    return LPSDynamicProgram(s)
}
