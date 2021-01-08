/*
Suppose you have n integers from 1 to n. We define a beautiful arrangement as an array that is constructed by these
n numbers successfully if one of the following is true for the ith position (1 <= i <= n) in this array:

The number at the ith position is divisible by i.
i is divisible by the number at the ith position.
Given an integer n, return the number of the beautiful arrangements that you can construct.

Example 1:

Input: n = 2
Output: 2
Explanation:
The first beautiful arrangement is [1, 2]:
Number at the 1st position (i=1) is 1, and 1 is divisible by i (i=1).
Number at the 2nd position (i=2) is 2, and 2 is divisible by i (i=2).
The second beautiful arrangement is [2, 1]:
Number at the 1st position (i=1) is 2, and 2 is divisible by i (i=1).
Number at the 2nd position (i=2) is 1, and i (i=2) is divisible by 1.
Example 2:

Input: n = 1
Output: 1
 Constraints:

1 <= n <= 15

*/
/*
Create the full array for every permutation and then check the array for the given divisibilty conditions.
Keep checking the elements while being added to the permutation array at every step for the divisibility condition and
stop creating it any further as soon as we find out the element just added to the permutation violates the divisiblity condition.
*/
func countArrangement(N int) int {
    if N == 0 {return 0}
    M := make([]int, N+1)
    for i:=0; i < N+1; i++ {M[i] = i}
    var helper func(s int)
    ret := 0
    helper = func(s int) {
        if s == 0 {
            ret++
            return
        }
        for i:=s; i > 0;i-- {
            M[i],M[s] = M[s], M[i]
            if M[s] % s == 0 || s % M[s] == 0 { helper(s-1)}
            M[i],M[s] = M[s], M[i]
        }
    }
    helper(N)
    return ret
}
