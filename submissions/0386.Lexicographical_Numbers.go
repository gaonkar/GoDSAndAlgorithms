/*
Given an integer n, return 1 - n in lexicographical order.

For example, given 13, return: [1,10,11,12,13,2,3,4,5,6,7,8,9].

Please optimize your algorithm to use less time and space. The input size may be as large as 5,000,000.
*/
func lexicalOrder(n int) (r []int) {
    stk := make([]int, 1)
    stk[0] = 1
    for len(stk) > 0 {
        //fmt.Println(stk)
        L := len(stk)
        x := stk[L-1]
        stk = stk[:L-1]
        if x + 1 <= n  && x %10 != 9 {
            stk = append(stk, x + 1)
        }
        if x * 10 <= n {
            stk = append(stk, x * 10)
        }
        r = append(r, x)
    }
    return r
}
