/*
Given a string S and a string T, count the number of distinct subsequences of S which equals T.

A subsequence of a string is a new string which is formed from the original string by deleting some (can be none) of
the characters without disturbing the relative positions of the remaining characters. (ie, "ACE" is a subsequence of
"ABCDE" while "AEC" is not).

It's guaranteed the answer fits on a 32-bit signed integer.

Example 1:

Input: S = "rabbbit", T = "rabbit"
Output: 3
Explanation:
As shown below, there are 3 ways you can generate "rabbit" from S.
(The caret symbol ^ means the chosen letters)

rabbbit
^^^^ ^^
rabbbit
^^ ^^^^
rabbbit
^^^ ^^^
Example 2:

Input: S = "babgbag", T = "bag"
Output: 5
Explanation:
As shown below, there are 5 ways you can generate "bag" from S.
(The caret symbol ^ means the chosen letters)

babgbag
^^ ^
babgbag
^^    ^
babgbag
^    ^^
babgbag
  ^  ^^
babgbag
    ^^^
*/


/*
 DP[i+1,j+1] --> number of distinct sub sequence from T[0..i] and S[0..j]
 DP[0, j+1]  --> 1 for all null T and any length S string
 DP[i+1, 0]  --> 0 for all null S string
 DP[i+1,j+1] --> DP[i+1, j] + D[i,j] && (T[i] == S[j])
                    If t[i] != s[j], the distinct subsequences will not include s[j] and thus all the number of distinct subsequences will simply be those in s[0..j - 1], which corresponds to dp[i][j]

 for j:=0; i < S; j++ {
    for i:=0; j < T; i++ {
        D[i+1][j+1] = DP[i+1][j] + int(T[i]==S[j]) * DP[i][j]
    }
 }
Each iteration, we only need D[i,j] so that can be simplified
 for j:=0; i < S; j++ {
    x := 1
    for i:=0; i < T; i++ {
        tmp := DP[i+1][j]
        DP[i+1][j+1] = DP[i+1][j] + int(T[i]==S[j]) * pre
        pre = tmp
    }
 }

 for j:=0; i < S; j++ {
    x := 1
    for i:=0; i < T; i++ {
        tmp := DP[i+1]
        DP[i+1] = DP[i+1] + int(T[i]==S[j]) * x
        pre = x
    }
 }

 if you loop in reverse of the inner loop, then you dont have to store the value for next iteration
  for j:=0; i < S; j++ {
    x := 1
    for i:=T; i > 0; i-- {
        DP[i+1] = DP[i+1] + int(T[i]==S[j]) * DP[i]
    }
 }

*/



func numDistinct(s string, t string) int {
    S := len(s)
    T := len(t)
    D := make([]int, T+1)
    D[0]=1
    for  j := 1; j <= S; j++ {
        for  i := T; i >= 1; i--{
            if s[j-1] == t[i-1] {
                D[i] += D[i-1]
            }
         }
        //fmt.Println(D[1:], s[:j])
     }
    return D[T];
}
