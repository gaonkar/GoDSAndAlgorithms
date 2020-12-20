/*
A string S of lowercase English letters is given. We want to partition this string into as many parts as 
possible so that each letter appears in at most one part, and return a list of integers representing the 
size of these parts.

 

Example 1:

Input: S = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.
 

Note:

S will have length in range [1, 500].
S will consist of lowercase English letters ('a' to 'z') only.
*/

func partitionLabels(S string) (r []int) {
    L := len(S)
    if L == 0 {return nil}
    lastpos := make([]int, 26)
    char := S[0]
    i:= L-1
    for i > 0 && S[i] != char{
        ch := int(S[i] - 'a')
        if lastpos[ch] == 0 {
            lastpos[ch] = i
        }
        i--
    }
    if i == 0 {
        r = append(r, 1)
        return append(r,partitionLabels(S[1:])...)
    }
    for j:=1; j < i;j++ {
        ch := int(S[j] - 'a')
        if lastpos[ch] == 0 {continue}
        if lastpos[ch] > i {i = lastpos[ch]}
    }
    r = append(r, i + 1)
    return append(r,partitionLabels(S[i+1:])...)
}