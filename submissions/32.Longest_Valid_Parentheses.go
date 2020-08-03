/*
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

Example 1:

Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
Example 2:

Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"


Track the valid brackets in a stack.
Invalid '(' tends to grow the stack.
Invalid ')' will make the stack goto 0, when that happens reset stk[0] to 0, but note the last longest valid bracket.
*/


func longestValidParentheses(s string) int {
    stk :=make([]int, 1)
    r,mr := 0,0
    for _,x:=range(s) {
        //fmt.Println(stk, r, mr, x)

        if x == '(' {
            stk = append(stk, 0)
            continue
        }
        L := len(stk)
        if L > 1 {
            mr = stk[L-1]
            mr +=2
            stk = stk[:L-1]
            stk[L-2] += mr
            continue
        }
        // saw too many ) note the last longest valid parentheses, and reset the value
        if r < stk[0] {r = stk[0]}
        stk[0] = 0
    }
    for _,x := range(stk) {
        if x > r {r = x}
    }
    return r
}

vector<vector<int>> moves = {{0,1},{1,0},{0,-1},{-1,0}}; //directions
        vector<vector<int>> res = {{r, c}};
        int x = 0, y = 1, tmp;
        int index = 0;
        for (int n = 0; res.size() < R * C; n++) {
            for (int i = 0; i < n / 2 + 1; i++) {
                r += moves[index%4][0], c += moves[index%4][1];
                if (0 <= r && r < R && 0 <= c && c < C)
                    res.push_back({r, c});
            }
            index++;
        }
        return res;
