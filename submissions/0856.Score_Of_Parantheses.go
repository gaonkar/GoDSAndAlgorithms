/*
Given a balanced parentheses string S, compute the score of the string based on the following rule:

() has score 1
AB has score A + B, where A and B are balanced parentheses strings.
(A) has score 2 * A, where A is a balanced parentheses string.


Example 1:

Input: "()"
Output: 1
Example 2:

Input: "(())"
Output: 2
Example 3:

Input: "()()"
Output: 2
Example 4:

Input: "(()(()))"
Output: 6


Note:

S is a balanced parentheses string, containing only ( and ).
2 <= S.length <= 50
*/
func scoreOfParentheses(S string) int {
    stk := make([]int, 0)
    for _,x := range(S) {
        //fmt.Println(stk, x)
        if x == '(' {
            stk = append(stk, -1)
            continue
        }
        if stk[len(stk)-1] == -1 {
            stk[len(stk)-1] = 1
            continue
        }
        tmp := 0
        for stk[len(stk)-1] != -1 {
            tmp += stk[len(stk)-1]
            stk = stk[:len(stk)-1]
        }
        stk[len(stk)-1] = tmp * 2
    }
    ret := 0
    for _,x := range(stk) {
        ret +=x
    }
    return ret
}
