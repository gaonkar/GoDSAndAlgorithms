/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Solution
1. All strings of length 1 are palindromes.
2. All strings of length 2 are palindromes if the next character is equal to itself.
3. For length l > 2, and position i  if s[i] == s[j+l], then it is part of palidrome if
    s[i+1].. s[j+l-1] is a palindrome.


O(N^2) time and O(3N) space

*/


func LPSDynamicProgram(s string) string {
    /*
    O(N^2) Space 4N
    */
    N := len(s)
    c:=0
    mx:=1
    var curr = make([]int, N)
    var prev1 = make([]int, N)
    var prev2 = make([]int, N)

    for i:= range prev2 {
        prev2[i] = 1
    }
    for i:= range prev1 {

        if i < N-1 && s[i] == s[i+1] {
            prev1[i] = 2
            if prev1[i] > mx {
                mx = prev1[i]
                c = i
            }
        }
    }
    for l:=3; l <= N; l++ { // length of the substring
        for i:= 0; i <= N-l;i++ {
            p:=i+l-1
            if s[i] == s[p] && prev2[i+1] > 0 {
                curr[i] = 2 + prev2[i+1]
            }
            if curr[i] > mx {
                mx = curr[i]
                c = i
            }
        }
        prev2= prev1
        prev1 = curr
        curr = make([]int, N)
        //fmt.Println(l, prev1)
    }
    //fmt.Println(c, mx, s[c:c+mx])

    return s[c:c+mx]
}

func longestPalindrome(s string) string {
    if len(s) <= 1 {
        return s
    }
    return LPSDynamicProgram(s)
}
