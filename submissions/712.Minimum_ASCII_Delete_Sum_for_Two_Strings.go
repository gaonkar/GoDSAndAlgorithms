/*
Given two strings s1, s2, find the lowest ASCII sum of deleted characters to make two strings equal.

Example 1:

Input: s1 = "sea", s2 = "eat"
Output: 231
Explanation: Deleting "s" from "sea" adds the ASCII value of "s" (115) to the sum.
Deleting "t" from "eat" adds 116 to the sum.
At the end, both strings are equal, and 115 + 116 = 231 is the minimum sum possible to achieve this.
Example 2:

Input: s1 = "delete", s2 = "leet"
Output: 403
Explanation: Deleting "dee" from "delete" to turn the string into "let",
adds 100[d]+101[e]+101[e] to the sum.  Deleting "e" from "leet" adds 101[e] to the sum.
At the end, both strings are equal to "let", and the answer is 100+101+101+101 = 403.
If instead we turned both strings into "lee" or "eet", we would get answers of 433 or 417, which are higher.
Note:

0 < s1.length, s2.length <= 1000.
All elements of each string will have an ASCII value in [97, 122].
*/

func Min(x, y int) int{
    if x < y {
        return x
    }
    return y
}

/*
 DP formulation
 DP[i,j] = Min delete between s[0:i] and t[0:j]
 DP[0,0] = 0
 DP[0,j] = DP[0,j-1] + t[j-1] for j > 0 // we have to delete all characters as one is null string
 DP[i,0] = DP[i-1,0] + s[i-1] for i > 0
 DP[i,j] = DP[i-1, j-1] if s[i-1] == t[j-1]
         = MIN(DP[i-1,j] + s[i-1], DP[i, j-1]+t[j-1])

*/

// Off by 1 as it makes it easier to understand
func MDS(text1 string, text2 string) int {
    var curr = make([]int, len(text1)+1)
    var prev = make([]int, len(text1)+1)
    //fmt.Println(curr, prev)
    for i:= 1; i <= len(text1); i++{
        prev[i] = prev[i-1] + int(text1[i-1]) // DP[0,j]
    }
    for j := range text2 {
        curr[0] =int(text2[0]) + prev[0] // DP[i,0]
        for i:= range text1{
            if text1[i] == text2[j] {
                curr[i+1] = prev[i]
            } else {
                curr[i+1] = Min(curr[i] + int(text1[i]), prev[i+1] + int(text2[j]))
            }
            fmt.Println(i,j, curr)
        }
        prev = curr
        curr = make([]int, len(text1)+1)
    }
    return prev[len(text1)]
}


func minimumDeleteSum(s1 string, s2 string) int {
    return MDS(s1, s2)
}

