/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

func GP(n int, o int, c int, M map[string] bool, t string) {
    if c == 0 { // all brackets are closed and we have n closed(c) brackets, add it to map
        M[t] = true
        return
    }
    if o < c { // you have an open bracket, so close it
        GP(n-1,o,c-1, M, t + ")")
    } 
    if o > 0 { // you have still brackets to open
        GP(n-1,o-1,c, M, t + "(")
    }
}

func generateParenthesis(n int) (r []string) {
    M := make(map[string] bool)
    GP(n,n,n, M, "")
    for k:= range(M) {
        r = append(r, k)
    }

    return r
}
