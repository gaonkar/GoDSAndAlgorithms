/*
Given an array of positive integers nums, remove the smallest subarray (possibly empty) such that the sum of the
remaining elements is divisible by p. It is not allowed to remove the whole array.

Return the length of the smallest subarray that you need to remove, or -1 if it's impossible.

A subarray is defined as a contiguous block of elements in the array.



Example 1:

Input: nums = [3,1,4,2], p = 6
Output: 1
Explanation: The sum of the elements in nums is 10, which is not divisible by 6. We can remove the subarray [4],
and the sum of the remaining elements is 6, which is divisible by 6.
Example 2:

Input: nums = [6,3,5,2], p = 9
Output: 2
Explanation: We cannot remove a single element to get a sum divisible by 9. The best way is to remove the subarray [5,2], leaving us with [6,3] with sum 9.
Example 3:

Input: nums = [1,2,3], p = 3
Output: 0
Explanation: Here the sum is 6. which is already divisible by 3. Thus we do not need to remove anything.
Example 4:

Input: nums = [1,2,3], p = 7
Output: -1
Explanation: There is no way to remove a subarray in order to get a sum divisible by 7.
Example 5:

Input: nums = [1000000000,1000000000,1000000000], p = 3
Output: 0


Constraints:

1 <= nums.length <= 105
1 <= nums[i] <= 109
1 <= p <= 109


[8,32,31,18,34,20,21,13,1,27,23,22,11,15,30,4,2]
148


*/
/*
    O(N^2)
    1. s:= sum(nums)
    2. P[i,j] = s - sum(nums[i..j])
    3. Now find the smallest i where S[i,...] is divisible
    We need to find (s - sum(i..j)) % K == 0
        S % K == sum(i..j) % K
 */


// C, C++, JAVA return negative MOD
func ModSum(a,b,K int) int {
    ret := (a+b) % K
    if ret < 0 {
        ret = (ret + K) % K
    }
    return ret
}

func minSubarray(A []int, K int) int {
    p,s := 0,0
    C := make(map[int] int)
    C[0] = -1
    for _,x := range(A) {
        s += x
    }
    s = ModSum(s,0,K)
    // no need to find any element
    if s == 0 {return 0}
    last := -1
    ret,ret2 := len(A), 100000
    for i,n:= range(A) {
        p = ModSum(p,n,K,)
        fmt.Println(i,s,p, ret,ret2, last, C)
        if s == p {     // I feel that we dont even need to track all the modulo, just the one that matches  S
                        // We have Off by 1 error.
            if ret2 > i - last {ret2 = i -last}
            last = i
        }
        idx := ModSum(p,-s,K)
        v, ok := C[idx]
        if ok && ret > (i-v){
            ret = i - v
        }
        C[p] = i // storing the index of the last time we saw modulo = i
    }
    fmt.Println(ret2,ret, len(A))
    if ret2 < len(A) {return ret2}
    return -1

    if ret <= len(A) {return ret}
    return -1
}
