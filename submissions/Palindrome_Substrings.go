/*
647
Given a string, your task is to count how many palindromic substrings in this string.

The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.

Example 1:

Input: "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".


Example 2:

Input: "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

Intuition
1. All substrings of len==1 is palindrome
2. All substrings of len==2 is palindrome if A[i] == A[i+1]
3. All substrings of len > 2 is palindrome if A[i] == A[j] iff A[i+1,j-1] is palindrome

*/

func countSubstrings(s string) int {
    N := len(s)
    C := make([]bool, N+1)
    P := C
    c := 0
    for i:=0; i < N; i++ {
        C = make([]bool, N+1)
        C[i] = true		// len==1 is palindrome
        c++
        for j:=i; j >=0; j-- { //upper diagonal start from [j,j] and go down to [j,0]
            if s[i]==s[j] {
                if i == j + 1 || P[j+1] { //condition 2 or 3
                    C[j] = true
                    c++
                }
            }
        }
        P = C	// copy the previous row
    }
    return c
}
