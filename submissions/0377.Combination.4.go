/*
Given an integer array with all positive numbers and no duplicates, find the number of possible combinations that add up to a positive integer target.

Example:

nums = [1, 2, 3]
target = 4

The possible combination ways are:
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)

Note that different sequences are counted as different combinations.

Therefore the output is 7.


Follow up:
What if negative numbers are allowed in the given array?
How does it change the problem?
What limitation we need to add to the question to allow negative numbers?

*/

// timed out
func CS4_recursive(N []int, t int, idx int) int {
    if t < 0 {return 0}
    if t == 0 {return 1}
    r := 0
    for i:=idx; i < len(N);i++  {
        r += CS4_recursive(N, t - N[i], 0)
    }
    return r
}


func CS4_DP(N []int, t int) int {
    D := make([]int, t+1)
    for s:= 1; s <= t; s++ {
        for i := 0; i < len(N); i++ {
            if s == N[i] {
                D[s]++
            } else if s >= N[i] {
                D[s] += D[s-N[i]]
            }
        }
    }
    return D[t]
}

func combinationSum4(nums []int, target int) int {
    return CS4_DP(nums, target)
}
