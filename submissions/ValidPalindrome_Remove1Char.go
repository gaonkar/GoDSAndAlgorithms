/*
Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.

Example 1:
Input: "aba"
Output: True
Example 2:
Input: "abca"
Output: True
Explanation: You could delete the character 'c'.
*/
func validPalindrome(s string) bool {
    i := 0
    j := len(s)-1
    skip := 1
    //fmt.Println(i,j)
    for i < j {
        if s[i] != s[j] && skip == 0 {
            break
        }
        if s[i] == s[j] {
            i++
            j--
        } else {
            skip = 0
            if s[i] == s[j-1] && s[i+1] == s[j] {
                ok := validPalindrome(s[i:j-1])
                if !ok {
                    ok = validPalindrome(s[i+1:j])
                }
                return ok
            }
            if s[i] == s[j-1] {
                j--
            } else if (s[i+1] == s[j]) {
                i++
            } else {
                break
            }
        }
        //fmt.Println(i,j, skip, s[:i], s[i:j], s[j:])
    }
    //fmt.Println(i,j, skip, s[:i], s[i:j], s[j:])
    return (i >= j)
}
