/*
Given an array A of integers, for each integer A[i] we need to choose either x = -K or x = K, and add x to A[i] (only once).

After this process, we have some array B.

Return the smallest possible difference between the maximum value of B and the minimum value of B.

 

Example 1:

Input: A = [1], K = 0
Output: 0
Explanation: B = [1]
Example 2:

Input: A = [0,10], K = 2
Output: 6
Explanation: B = [2,8]
Example 3:

Input: A = [1,3,6], K = 3
Output: 3
Explanation: B = [4,6,3]
 

Note:

1 <= A.length <= 10000
0 <= A[i] <= 10000
0 <= K <= 10000
*/
func max(x,y int) int {
    if x < y {return y}
    return x
}

func min(x,y int) int {
    if x < y {return x}
    return y
}

/*
1. {+K,-K} --> {0, -2*K}
2. We subtract from i to N-1 by 2K 
    A0........Ai-1,Ai-2K......AN-2K
    case Ai-1 > AN-2K we have to choose Ai-1, so nMax = Max(Ai-1, AN-2K)
    case Ai - 2K < A0, we have to choose Ai-2K, so nMin = Min(Ai-2K, A0)
    now we have to iterate from 1 to N-1 to find all such splits
 3. We can do more optimizations where we only loop until Ai -2K < A0 and from A[i] > AN-2K, 
    thats value dependent
*/

func smallestRangeII(A []int, K int) int {
    sort.Ints(A)

    A0, AN := A[0], A[len(A)-1]
    diff:= AN - A0
    for i:=1; i < len(A); i++ {
        nMax, nMin := max(AN-2*K, A[i-1]), min(A0, A[i]-2*K)
        diff = min(diff, nMax-nMin)
    }
    return diff
}t
