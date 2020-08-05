/*
You are given a list of non-negative integers, a1, a2, ..., an, and a target, S.
Now you have 2 symbols + and -. For each integer, you should choose one from + and - as its new symbol.

Find out how many ways to assign symbols to make sum of integers equal to target S.

Example 1:

Input: nums is [1, 1, 1, 1, 1], S is 3.
Output: 5

Now after assigning them + and  - sign, we get positives and negatives set of numbers
sum(nums) = S
sum(positives) - sum(negatives) = T
adding both, we get
sum(positives) = (S+T) / 2

This is a subset sum problem with target T

*/

func SubsetSumCount(nums []int, T int) int {
    DP:= make([]int, T+1)
    DP[0] = 1
    for i:= range(nums) {
        for j := T; j >= nums[i]; j-- {
            DP[j] += DP[j-nums[i]]
        }
    }
    return DP[T]
}

func findTargetSumWays(nums []int, S int) int {
    sum := 0
    for _,x := range(nums) {sum += x}
    if sum < S || (sum + S) % 2 == 0 {return 0} // sum +S should be even
    return SubsetSumCount(nums, (sum+S)/2)
}


// solution with Memoization

func findTargetSumWays(nums []int, S int) (r int) {
    var helper func(idx, cs int) int
    L := len(nums)
    DP := make([]map[int]int,L)
    for i:=0; i < L;i++ {
        DP[i] = map[int] int{}
    }
    //sort.Ints(nums)
    sum := 0
    for i:=0; i < len(nums);i++ { sum += nums[i]}
    if sum < S {return 0}
    helper = func(idx, cs int) int {
        if L == idx {
            if cs == 0 {return 1}
            return 0
        } else if DP[idx][cs] > 0 {
            return DP[idx][cs]
        }

        l := helper(idx+1, cs+nums[idx])
        r := helper(idx+1, cs-nums[idx])
        DP[idx][cs] = l + r
        return l+r
    }
    return helper(0,S)
}
