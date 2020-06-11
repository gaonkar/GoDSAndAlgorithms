/*
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated
exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.

Furthermore, you may assume that the original data does not contain any digits and that digits are only for those
repeat numbers, k. For example, there won't be input like 3a or 2[4].



Example 1:

Input: s = "3[a]2[bc]"
Output: "aaabcbc"
Example 2:

Input: s = "3[a2[c]]"
Output: "accaccacc"
Example 3:

Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"
Example 4:

Input: s = "abc3[cd]xyz"
Output: "abccdcdcdxyz"


Painful problem
*/

func IsDigit(x byte) bool {
    if '0' <= x && x <= '9' {return true}
    return false
}
func IsBracket(x byte) bool {
    if '[' == x || x == ']' {return true}
    return false
}



func DT(s string, i int) (r string, rI int) {
    if len(s) <= i || s[i] == ']' {return "", i}
    //fmt.Println(s[i:], i)
    t := ""
    for i < len(s) && s[i] != ']' {
        //fmt.Println("B", i, r)
        if !IsDigit(s[i]) {
            r = r + string(s[i])
            i++
            continue
        }
        N:=0
        for i < len(s) && IsDigit(s[i]){
            N = N * 10 + int(s[i]-'0')
            i++
        }
        i++
        t, i = DT(s,i) // bug was here, t,i := reset the value of i, need to look into implementation
        i++
        //fmt.Println("T", r, t, N, i, s[i:])
        for N > 0 {
            r = r + t
            N--
        }
        //fmt.Println("N:", i, r)
    }
    //fmt.Println("R", r, i)
    return r, i
}

func decodeString(s string) (r string) {
    i := 0
    r, i  = DT(s, i)
    return r
}
