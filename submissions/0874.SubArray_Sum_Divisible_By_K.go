/*
Given an array A of integers, return the number of (contiguous, non-empty) subarrays that have a sum divisible by K.



Example 1:

Input: A = [4,5,0,-2,-3,1], K = 5
Output: 7
Explanation: There are 7 subarrays with a sum divisible by K = 5:
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]


Note:

1 <= A.length <= 30000
-10000 <= A[i] <= 10000
2 <= K <= 10000
*/



/*
    Find all A[i..j] whose sum(A[i..j]) %K == 0
    (P[j] - P[i]) %K == 0
    P[j]%K == P[i] % K
    // count all prefix that go to same modulo, then add them up
 */

func subarraysDivByK(A []int, K int) int {
    c:= 0
    s:= 0
    C := make([]int, K+1)
    C[0] = 1
    for _,n:= range(A) {
        s = (K+ s + n%K) % K
        c += C[s] // simple counting 1 + 2 + 3 + 4 ... so on 
        C[s]++
    }
    return c
}
