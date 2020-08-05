/*
Given a non-empty string check if it can be constructed by taking a substring of it and appending multiple copies
of the substring together. You may assume the given string consists of lowercase English letters only and its length
will not exceed 10000.



Example 1:

Input: "abab"
Output: True
Explanation: It's the substring "ab" twice.
Example 2:

Input: "aba"
Output: False
Example 3:

Input: "abcabcabcabc"
Output: True
Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)

If we assume that the string has a repeating pattern, then A[0] must be first character and A[N] must be last character
Take "abccabcc",
Now if we concatenate s to itself, "abccabccabccabcc" then delete first and last character, we get "bccabccabccabc".
search the original string, and it should be found at first position at repeat.


*/

func repeatedSubstringPattern(s string) bool {
    r := s + s
    r = r[1:]
    r = r[:len(r)-1]
    return strings.Contains(r,s)
}
