/*
Given an array of integers nums and a positive integer k, find whether it's possible to divide this array into k non-empty subsets whose sums are all equal.



Example 1:

Input: nums = [4, 3, 2, 3, 5, 2, 1], k = 4
Output: True
Explanation: It's possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.


Note:

1 <= k <= len(nums) <= 16.
0 < nums[i] < 10000.
*/
func CPK(N []int, V []bool, k int, S int, tgt int, idx int) bool {
    if k == 1 { return true}
    if S == 0 {
        //fmt.Println(k, V)
        return CPK(N, V, k-1, tgt,tgt, 0)
    }
    for i:= idx; i < len(N); i++ {
        if !V[i] {
            V[i] = true
            if CPK(N, V, k, S-N[i],tgt, i+1) == true {return true}
            V[i] = false
        }
    }
    return false
}

func canPartitionKSubsets(nums []int, k int) bool {
    S:=0
    for _,x := range(nums) { S+=x}
    if S % k != 0 {return false}
    S = S/k
    return CPK(nums, make([]bool, len(nums)), k, S,S, 0)
}
