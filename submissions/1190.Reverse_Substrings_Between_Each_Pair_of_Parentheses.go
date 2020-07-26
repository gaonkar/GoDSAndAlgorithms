/*
1190.Reverse_Substrings_Between_Each_Pair_of_Parentheses.go

You are given a string s that consists of lower case English letters and brackets.

Reverse the strings in each pair of matching parentheses, starting from the innermost one.

Your result should not contain any brackets.



Example 1:

Input: s = "(abcd)"
Output: "dcba"
Example 2:

Input: s = "(u(love)i)"
Output: "iloveu"
Explanation: The substring "love" is reversed first, then the whole string is reversed.
Example 3:

Input: s = "(ed(et(oc))el)"
Output: "leetcode"
Explanation: First, we reverse the substring "oc", then "etco", and finally, the whole string.
Example 4:

Input: s = "a(bcdefghijkl(mno)p)q"
Output: "apmnolkjihgfedcbq"

*/

// O(# of brackets)

func Reverse(runes []rune, i, j int) {
    for i < j{
        runes[i], runes[j] = runes[j], runes[i]
        i++
        j--
    }
}

func reverseParentheses(s string) string {
    runes := []rune(s)
    stk:= make([]int, len(s))
    stkLen := 0
    for i,x := range(s) {
        if x == '(' {
            pIdx := i+1
            for pIdx < len(s) && s[pIdx] == '(' {pIdx++}
            stk[stkLen] = pIdx
            stkLen++
            continue
        }
        if x == ')' {
            stkLen--
            pIdx := stk[stkLen]
            Reverse(runes,pIdx, i-1)
            continue
        }
    }
    ret:= []rune("")

    for _, x := range(runes) {
        if x != '(' && x != ')' {ret = append(ret, x)}
    }
    return string(ret)

}

// O(N) From lee215

func reverseParentheses(s string) string {
    p := make(map[int]int)
    stk:= make([]int, len(s))
    stkLen := 0
    for i,x := range(s) {
        if x == '(' {
            stk[stkLen] = i
            stkLen++
            continue
        }
        if x == ')' {
            stkLen--
            pIdx := stk[stkLen]
            p[i] = pIdx
            p[pIdx] = i
            continue
        }
    }
    i, d, c := 0, 1, 0
    ret := make([]byte, 0)
    for c < len(s) {
        // we need to reverse whenever we see a ()
        /*
            Example "ab(cd)ef"
                     01234567
            ret copies upto ab and then sees (, we set i as 5 and dir -1, i is decremented to 4
            we copy dc, then we see ( again, but now set i = 5 and dir 1, i is incremented to 6.
            we copy ef, giving us abdcef 
        */
        if s[i] == ')' || s[i] == '(' {
            i = p[i]
            d = -d
        } else {
            ret = append(ret, s[i])
        }
        i +=d
        c++
    }
    return string(ret)
}
