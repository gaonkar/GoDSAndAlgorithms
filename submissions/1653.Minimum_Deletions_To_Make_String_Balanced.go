/*
You are given a string s consisting only of characters 'a' and 'b'​​​​.
You can delete any number of characters in s to make s balanced. s is balanced if there is no
pair of indices (i,j) such that i < j and s[i] = 'b' and s[j]= 'a'.
Return the minimum number of deletions needed to make s balanced.



Example 1:

Input: s = "aababbab"
Output: 2
Explanation: You can either:
Delete the characters at 0-indexed positions 2 and 6 ("aababbab" -> "aaabbb"), or
Delete the characters at 0-indexed positions 3 and 6 ("aababbab" -> "aabbbb").
Example 2:

Input: s = "bbaaaaabb"
Output: 2
Explanation: The only solution is to delete the first two characters.


Constraints:

1 <= s.length <= 105
s[i] is 'a' or 'b'

Solution
find the location i,i+1 such that number of b seen upto i + number of a seen from len(s) to i+1 is minimum

*/

func minimumDeletions(s string) int {
    A := make([]int, len(s)+1)
    B := make([]int, len(s)+1)
    a, b:=0,0
    for i,_ := range(s) {
        j := len(s) - i - 1
        A[j], B[i+1] = a,b
        if s[i] == 'b' {
            b++
            B[i+1] =b
        }
        if s[j] == 'a' {
            a++
            A[j] = a
        }
    }
    min := len(s)
    for i := 0; i < len(A); i++ {
        if min > A[i] + B[i] {
            min = A[i] + B[i]
        }
    }
    return min
}
